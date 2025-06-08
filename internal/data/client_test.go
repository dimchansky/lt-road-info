package data

import (
	"testing"
)

func TestClient_FetchEALData(t *testing.T) {
	// For now, this is a basic integration test
	// In the future, we can add VCR recording
	client := NewClient(nil)
	
	layers, err := client.FetchEALData()
	if err != nil {
		t.Fatalf("Failed to fetch EAL data: %v", err)
	}
	
	if len(layers) == 0 {
		t.Error("Expected at least one layer, got none")
	}
	
	// Basic validation
	for _, layer := range layers {
		if layer.Layer == "" {
			t.Error("Layer should have a name")
		}
		if len(layer.Features) == 0 {
			t.Logf("Layer %s has no features", layer.Layer)
		}
	}
}

func TestClient_FetchArcGISData(t *testing.T) {
	// Basic integration test
	client := NewClient(nil)
	
	features, err := client.FetchArcGISData()
	if err != nil {
		t.Fatalf("Failed to fetch ArcGIS data: %v", err)
	}
	
	if len(features) == 0 {
		t.Error("Expected at least one feature, got none")
	}
	
	// Basic validation
	for i, feature := range features {
		if len(feature.Geometry.Paths) == 0 {
			t.Errorf("Feature %d should have at least one path", i)
		}
		if feature.Attributes == nil {
			t.Errorf("Feature %d should have attributes", i)
		}
		
		// Just test first few features to avoid long tests
		if i >= 5 {
			break
		}
	}
}