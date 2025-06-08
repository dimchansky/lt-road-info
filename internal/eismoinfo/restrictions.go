package eismoinfo

import (
	"net/http"

	"github.com/dimchansky/lt-road-info/internal/converter"
	"github.com/dimchansky/lt-road-info/internal/data"
)

// DownloadRestrictions downloads road restrictions and saves them as GPX
func DownloadRestrictions(outputPath string) error {
	return DownloadRestrictionsWithClient(http.DefaultClient, outputPath)
}

// DownloadRestrictionsWithClient downloads restrictions using a custom HTTP client
// This allows for testing with go-vcr or other HTTP interceptors
func DownloadRestrictionsWithClient(httpClient *http.Client, outputPath string) error {
	// Create data client
	client := data.NewClient(httpClient)

	// Fetch data
	layers, err := client.FetchEALData()
	if err != nil {
		return err
	}

	// Convert to GPX
	return converter.EALToGPX(layers, outputPath)
}
