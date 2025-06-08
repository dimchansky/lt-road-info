package eismoinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dimchansky/lt-road-info/internal/transform"
	"github.com/tkrajina/gpxgo/gpx"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/charmap"
)

const restrictionsURL = "https://eismoinfo.lt/eismoinfo-backend/layer-dynamic-features/EAL?lks=true"

// Restriction represents a road restriction from the EAL API
type Restriction struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type Properties struct {
	LayerIcon  int    `json:"layerIcon"`
	Tooltip    string `json:"tooltip"`
	UpdateDate string `json:"updateDate"`
}

type Geometry struct {
	Type        string          `json:"type"`
	Coordinates json.RawMessage `json:"coordinates"`
}

// DownloadRestrictions downloads road restrictions and saves them as GPX
func DownloadRestrictions(outputPath string) error {
	// Fetch data from API
	resp, err := http.Get(restrictionsURL)
	if err != nil {
		return fmt.Errorf("failed to fetch restrictions: %w", err)
	}
	defer resp.Body.Close()

	// Handle character encoding
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		reader = resp.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// Parse JSON
	var data Restriction
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Create GPX
	gpxData := gpx.GPX{
		Version: "1.1",
		Creator: "lt-road-info",
		Name:    "Lithuanian Road Restrictions",
		Time:    &time.Time{},
	}
	*gpxData.Time = time.Now()

	// Process features
	for _, feature := range data.Features {
		track := gpx.GPXTrack{
			Name: fmt.Sprintf("%s - %s", getRestrictionName(feature.Properties.LayerIcon), cleanDescription(feature.Properties.Tooltip)),
		}

		segment := gpx.GPXTrackSegment{}

		// Parse coordinates based on geometry type
		if err := parseGeometry(feature.Geometry, &segment); err != nil {
			continue // Skip invalid geometries
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

func parseGeometry(geom Geometry, segment *gpx.GPXTrackSegment) error {
	switch geom.Type {
	case "LineString":
		var coords [][]float64
		if err := json.Unmarshal(geom.Coordinates, &coords); err != nil {
			return err
		}
		for _, coord := range coords {
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
	case "MultiLineString":
		var multiCoords [][][]float64
		if err := json.Unmarshal(geom.Coordinates, &multiCoords); err != nil {
			return err
		}
		for _, line := range multiCoords {
			for _, coord := range line {
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
	}
	return nil
}

func getRestrictionName(layerIcon int) string {
	names := map[int]string{
		1:  "Road Closed",
		2:  "Road Works",
		3:  "Slippery Road",
		4:  "Traffic Accident",
		5:  "Speed Limit",
		6:  "Weight Limit",
		7:  "Width Limit",
		8:  "Height Limit",
		9:  "Dangerous Section",
		10: "Other Restriction",
	}
	
	if name, ok := names[layerIcon]; ok {
		return name
	}
	return fmt.Sprintf("Restriction Type %d", layerIcon)
}

func cleanDescription(desc string) string {
	// Convert Windows-1257 to UTF-8 if needed
	decoder := charmap.Windows1257.NewDecoder()
	utf8Desc, _ := decoder.String(desc)
	
	// Clean up HTML entities and extra spaces
	utf8Desc = strings.ReplaceAll(utf8Desc, "&nbsp;", " ")
	utf8Desc = strings.ReplaceAll(utf8Desc, "  ", " ")
	utf8Desc = strings.TrimSpace(utf8Desc)
	
	return utf8Desc
}