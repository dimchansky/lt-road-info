package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dimchansky/lt-road-info/internal/arcgis"
	"github.com/dimchansky/lt-road-info/internal/eismoinfo"
	"github.com/tkrajina/gpxgo/gpx"
)

func main() {
	fmt.Println("🔍 Verifying coordinate transformations...")

	// Create temporary directory
	tmpDir := "/tmp/lt-road-verify"
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir)

	// Test restrictions download
	fmt.Println("\n📍 Testing road restrictions...")
	restrictionsPath := filepath.Join(tmpDir, "test-restrictions.gpx")
	err := eismoinfo.DownloadRestrictions(restrictionsPath)
	if err != nil {
		fmt.Printf("❌ Failed to download restrictions: %v\n", err)
		os.Exit(1)
	}

	valid, err := validateGPXCoordinates(restrictionsPath, "restrictions")
	if err != nil {
		fmt.Printf("❌ Error validating restrictions: %v\n", err)
		os.Exit(1)
	}
	if !valid {
		fmt.Println("❌ Restrictions validation failed")
		os.Exit(1)
	}

	// Test speed control download
	fmt.Println("\n🚗 Testing speed control sections...")
	speedPath := filepath.Join(tmpDir, "test-speed.gpx")
	err = arcgis.DownloadSpeedControlSections(speedPath)
	if err != nil {
		fmt.Printf("❌ Failed to download speed control: %v\n", err)
		os.Exit(1)
	}

	valid, err = validateGPXCoordinates(speedPath, "speed control")
	if err != nil {
		fmt.Printf("❌ Error validating speed control: %v\n", err)
		os.Exit(1)
	}
	if !valid {
		fmt.Println("❌ Speed control validation failed")
		os.Exit(1)
	}

	fmt.Println("\n✅ All coordinate transformations are correct!")
	fmt.Println("🇱🇹 All GPX tracks are properly located in Lithuania")
}

func validateGPXCoordinates(filePath, trackType string) (bool, error) {
	gpxFile, err := gpx.ParseFile(filePath)
	if err != nil {
		return false, fmt.Errorf("failed to parse GPX: %w", err)
	}

	if len(gpxFile.Tracks) == 0 {
		return false, fmt.Errorf("no tracks found in GPX file")
	}

	validCount := 0
	totalCount := 0
	var sampleCoords []string

	for _, track := range gpxFile.Tracks {
		for _, segment := range track.Segments {
			for _, point := range segment.Points {
				totalCount++

				// Check if coordinates are in Lithuania
				if isInLithuania(point.Latitude, point.Longitude) {
					validCount++

					// Collect some sample coordinates
					if len(sampleCoords) < 3 {
						sampleCoords = append(sampleCoords,
							fmt.Sprintf("[%.6f, %.6f]", point.Latitude, point.Longitude))
					}
				} else {
					// Check if they're in the Abu Dhabi area (our previous bug)
					if isInAbuDhabiArea(point.Latitude, point.Longitude) {
						fmt.Printf("❌ %s coordinate [%.6f, %.6f] is in Abu Dhabi - lat/lon mixup detected!\n",
							trackType, point.Latitude, point.Longitude)
						return false, nil
					}

					fmt.Printf("⚠️  %s coordinate [%.6f, %.6f] is outside Lithuania\n",
						trackType, point.Latitude, point.Longitude)
				}
			}
		}
	}

	validPercent := float64(validCount) / float64(totalCount) * 100

	fmt.Printf("   📊 %s: %d/%d coordinates in Lithuania (%.1f%%)\n",
		trackType, validCount, totalCount, validPercent)

	if len(sampleCoords) > 0 {
		fmt.Printf("   📍 Sample coordinates: %v\n", sampleCoords)
	}

	// Require at least 90% of coordinates to be in Lithuania
	if validPercent < 90.0 {
		fmt.Printf("❌ Only %.1f%% of coordinates are in Lithuania (expected >90%%)\n", validPercent)
		return false, nil
	}

	fmt.Printf("✅ %s coordinates validation passed\n", trackType)
	return true, nil
}

func isInLithuania(lat, lon float64) bool {
	// Lithuania approximate boundaries
	return lat >= 53.5 && lat <= 56.5 && lon >= 20.5 && lon <= 27.0
}

func isInAbuDhabiArea(lat, lon float64) bool {
	// Abu Dhabi approximate area (where our bug put coordinates)
	return lat >= 24.0 && lat <= 25.0 && lon >= 54.0 && lon <= 56.0
}
