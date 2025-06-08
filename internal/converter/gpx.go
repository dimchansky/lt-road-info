package converter

import (
	"fmt"
	"os"
	"time"

	"github.com/dimchansky/lt-road-info/internal/data"
	"github.com/dimchansky/lt-road-info/internal/transform"
	"github.com/tkrajina/gpxgo/gpx"
)

// EALToGPX converts EAL data to GPX format and saves to file
func EALToGPX(layers []data.EALLayer, outputPath string) error {
	// Create GPX
	gpxData := gpx.GPX{
		Version: "1.1",
		Creator: "lt-road-info",
		Name:    "Lithuanian Road Restrictions",
		Time:    &time.Time{},
	}
	*gpxData.Time = time.Now()

	// Process all features from all layers
	for _, layer := range layers {
		for _, feature := range layer.Features {
			// Process each restriction within the feature
			for _, restriction := range feature.Restrictions {
				track := gpx.GPXTrack{
					Name: fmt.Sprintf("%s - %s", feature.Name, getRestrictionDescription(restriction)),
				}

				// Process all paths in the restriction
				for _, path := range restriction.Lines.Paths {
					segment := gpx.GPXTrackSegment{}
					
					// Convert coordinates to GPX points
					for _, coord := range path {
						if len(coord) >= 2 {
							lat, lon := transform.LKS94ToWGS84(coord[0], coord[1])
							segment.Points = append(segment.Points, gpx.GPXPoint{
								Point: gpx.Point{
									Latitude:  lat,
									Longitude: lon,
								},
							})
						}
					}
					
					if len(segment.Points) > 0 {
						track.Segments = append(track.Segments, segment)
					}
				}
				
				if len(track.Segments) > 0 {
					gpxData.Tracks = append(gpxData.Tracks, track)
				}
			}
		}
	}

	// Save to file
	return saveGPX(gpxData, outputPath)
}

// ArcGISToGPX converts ArcGIS speed control data to GPX format
func ArcGISToGPX(features []data.ArcGISFeature, outputPath string) error {
	// Create GPX
	gpxData := gpx.GPX{
		Version: "1.1",
		Creator: "lt-road-info",
		Name:    "Lithuanian Speed Control Sections",
		Time:    &time.Time{},
	}
	*gpxData.Time = time.Now()

	// Process features
	for i, feature := range features {
		track := gpx.GPXTrack{
			Name: fmt.Sprintf("Speed Control Section %d", i+1),
		}

		// Add description if available
		if desc := getArcGISFeatureDescription(feature); desc != "" {
			track.Name = fmt.Sprintf("%s - %s", track.Name, desc)
		}

		// Process geometry paths
		for _, path := range feature.Geometry.Paths {
			segment := gpx.GPXTrackSegment{}
			
			for _, coord := range path {
				if len(coord) >= 2 {
					lat, lon := transform.LKS94ToWGS84(coord[0], coord[1])
					segment.Points = append(segment.Points, gpx.GPXPoint{
						Point: gpx.Point{
							Latitude:  lat,
							Longitude: lon,
						},
					})
				}
			}
			
			if len(segment.Points) > 0 {
				track.Segments = append(track.Segments, segment)
			}
		}

		if len(track.Segments) > 0 {
			gpxData.Tracks = append(gpxData.Tracks, track)
		}
	}

	// Save to file
	return saveGPX(gpxData, outputPath)
}

func saveGPX(gpxData gpx.GPX, outputPath string) error {
	xmlBytes, err := gpxData.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})
	if err != nil {
		return fmt.Errorf("failed to generate GPX XML: %w", err)
	}

	if err := os.WriteFile(outputPath, xmlBytes, 0644); err != nil {
		return fmt.Errorf("failed to write GPX file: %w", err)
	}

	return nil
}

func getRestrictionDescription(restriction data.EALRestriction) string {
	desc := fmt.Sprintf("Restriction %s", restriction.Icon)
	if restriction.IconValue > 0 {
		desc += fmt.Sprintf(" (%.0f)", restriction.IconValue)
	}
	return desc
}

func getArcGISFeatureDescription(feature data.ArcGISFeature) string {
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