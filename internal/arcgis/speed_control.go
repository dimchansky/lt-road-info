package arcgis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/dimchansky/lt-road-info/internal/transform"
	"github.com/tkrajina/gpxgo/gpx"
)

const (
	baseURL      = "https://gis.ktvis.lt/arcgis/rest/services/PUB/PUB_ITS/MapServer/13"
	queryURL     = baseURL + "/query"
	maxBatchSize = 1000
)

// Feature represents a speed control section from ArcGIS
type Feature struct {
	Attributes map[string]interface{} `json:"attributes"`
	Geometry   Geometry               `json:"geometry"`
}

type Geometry struct {
	Paths [][][]float64 `json:"paths"`
}

type QueryResponse struct {
	Features         []Feature `json:"features"`
	ExceededTransfer bool      `json:"exceededTransferLimit"`
}

type ServiceInfo struct {
	MaxRecordCount int `json:"maxRecordCount"`
}

// DownloadSpeedControlSections downloads speed control sections and saves them as GPX
func DownloadSpeedControlSections(outputPath string) error {
	// Get service information
	maxRecords, err := getMaxRecordCount()
	if err != nil {
		return fmt.Errorf("failed to get service info: %w", err)
	}

	// Create GPX
	gpxData := gpx.GPX{
		Version: "1.1",
		Creator: "lt-road-info",
		Name:    "Lithuanian Speed Control Sections",
		Time:    &time.Time{},
	}
	*gpxData.Time = time.Now()

	// Fetch all features
	allFeatures, err := fetchAllFeatures(maxRecords)
	if err != nil {
		return fmt.Errorf("failed to fetch features: %w", err)
	}

	// Process features
	for i, feature := range allFeatures {
		track := gpx.GPXTrack{
			Name: fmt.Sprintf("Speed Control Section %d", i+1),
		}

		// Add description if available
		if desc := getFeatureDescription(feature); desc != "" {
			track.Name = fmt.Sprintf("%s - %s", track.Name, desc)
		}

		segment := gpx.GPXTrackSegment{}

		// Process geometry paths
		for _, path := range feature.Geometry.Paths {
			for _, coord := range path {
				if len(coord) >= 2 {
					lon, lat := transform.LKS94ToWGS84(coord[0], coord[1])
					segment.Points = append(segment.Points, gpx.GPXPoint{
						Point: gpx.Point{
							Latitude:  lat,
							Longitude: lon,
						},
					})
				}
			}
		}

		if len(segment.Points) > 0 {
			track.Segments = append(track.Segments, segment)
			gpxData.Tracks = append(gpxData.Tracks, track)
		}
	}

	// Save to file
	xmlBytes, err := gpxData.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})
	if err != nil {
		return fmt.Errorf("failed to generate GPX XML: %w", err)
	}

	if err := os.WriteFile(outputPath, xmlBytes, 0644); err != nil {
		return fmt.Errorf("failed to write GPX file: %w", err)
	}

	return nil
}

func getMaxRecordCount() (int, error) {
	resp, err := http.Get(baseURL + "?f=json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var info ServiceInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return 0, err
	}

	if info.MaxRecordCount > 0 {
		return info.MaxRecordCount, nil
	}
	return maxBatchSize, nil
}

func fetchAllFeatures(maxRecords int) ([]Feature, error) {
	var allFeatures []Feature
	offset := 0

	for {
		features, hasMore, err := fetchFeatureBatch(offset, maxRecords)
		if err != nil {
			return nil, err
		}

		allFeatures = append(allFeatures, features...)

		if !hasMore {
			break
		}

		offset += len(features)
	}

	return allFeatures, nil
}

func fetchFeatureBatch(offset, limit int) ([]Feature, bool, error) {
	params := url.Values{
		"where":          {"1=1"},
		"outFields":      {"*"},
		"returnGeometry": {"true"},
		"f":              {"json"},
		"resultOffset":   {strconv.Itoa(offset)},
		"resultRecordCount": {strconv.Itoa(limit)},
		"outSR":          {"3346"}, // LKS-94
	}

	resp, err := http.Get(queryURL + "?" + params.Encode())
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}

	var result QueryResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, false, err
	}

	return result.Features, result.ExceededTransfer, nil
}

func getFeatureDescription(feature Feature) string {
	var desc string
	
	// Extract relevant attributes
	if roadName, ok := feature.Attributes["road_name"].(string); ok && roadName != "" {
		desc = roadName
	}
	
	if roadNum, ok := feature.Attributes["road_number"].(string); ok && roadNum != "" {
		if desc != "" {
			desc += " (" + roadNum + ")"
		} else {
			desc = "Road " + roadNum
		}
	}
	
	if speedLimit, ok := feature.Attributes["speed_limit"]; ok {
		if desc != "" {
			desc += " - "
		}
		desc += fmt.Sprintf("Speed limit: %v km/h", speedLimit)
	}
	
	return desc
}