package converter

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dimchansky/lt-road-info/internal/data"
)

func TestEALToGPX(t *testing.T) {
	// Create test data
	testLayers := []data.EALLayer{
		{
			Layer: "EAL",
			Name:  "Test Layer",
			Features: []data.EALFeature{
				{
					ID:   "test-1",
					Name: "Test Restriction",
					Restrictions: []data.EALRestriction{
						{
							ID:        "restriction-1",
							Icon:      "76",
							IconValue: 50.0,
							Lines: data.EALLines{
								Paths: [][][]float64{
									{
										{532186, 6190040},
										{532189, 6190044},
										{532218, 6190080},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create temporary file
	tmpDir := t.TempDir()
	outputPath := filepath.Join(tmpDir, "test-restrictions.gpx")

	// Convert to GPX
	err := EALToGPX(testLayers, outputPath)
	if err != nil {
		t.Fatalf("Failed to convert EAL to GPX: %v", err)
	}

	// Check file exists and has content
	info, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf("Output file not created: %v", err)
	}

	if info.Size() == 0 {
		t.Error("Output file is empty")
	}

	// Read and basic validate content
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	contentStr := string(content)
	if !contains(contentStr, "<?xml") {
		t.Error("Output should be XML")
	}
	if !contains(contentStr, "Test Restriction") {
		t.Error("Output should contain test data")
	}
}

func TestArcGISToGPX(t *testing.T) {
	// Create test data
	testFeatures := []data.ArcGISFeature{
		{
			Attributes: map[string]interface{}{
				"road_name":   "Test Road",
				"speed_limit": 90,
			},
			Geometry: data.ArcGISGeometry{
				Paths: [][][]float64{
					{
						{532186, 6190040},
						{532189, 6190044},
						{532218, 6190080},
					},
				},
			},
		},
	}

	// Create temporary file
	tmpDir := t.TempDir()
	outputPath := filepath.Join(tmpDir, "test-speed-control.gpx")

	// Convert to GPX
	err := ArcGISToGPX(testFeatures, outputPath)
	if err != nil {
		t.Fatalf("Failed to convert ArcGIS to GPX: %v", err)
	}

	// Check file exists and has content
	info, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf("Output file not created: %v", err)
	}

	if info.Size() == 0 {
		t.Error("Output file is empty")
	}

	// Read and basic validate content
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	contentStr := string(content)
	if !contains(contentStr, "<?xml") {
		t.Error("Output should be XML")
	}
	if !contains(contentStr, "Test Road") {
		t.Error("Output should contain test data")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsInMiddle(s, substr))
}

func containsInMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}