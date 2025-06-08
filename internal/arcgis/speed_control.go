package arcgis

import (
	"net/http"

	"github.com/dimchansky/lt-road-info/internal/converter"
	"github.com/dimchansky/lt-road-info/internal/data"
)

// DownloadSpeedControlSections downloads speed control sections and saves them as GPX
func DownloadSpeedControlSections(outputPath string) error {
	return DownloadSpeedControlSectionsWithClient(http.DefaultClient, outputPath)
}

// DownloadSpeedControlSectionsWithClient downloads speed control sections using a custom HTTP client
// This allows for testing with go-vcr or other HTTP interceptors
func DownloadSpeedControlSectionsWithClient(httpClient *http.Client, outputPath string) error {
	// Create data client
	client := data.NewClient(httpClient)
	
	// Fetch data
	features, err := client.FetchArcGISData()
	if err != nil {
		return err
	}
	
	// Convert to GPX
	return converter.ArcGISToGPX(features, outputPath)
}