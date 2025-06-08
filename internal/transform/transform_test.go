package transform

import (
	"testing"
)

func TestLKS94ToWGS84CoordinateOrder(t *testing.T) {
	// This test specifically prevents the lat/lon coordinate mixup bug
	// that caused tracks to appear in Abu Dhabi instead of Lithuania
	
	testCases := []struct {
		name          string
		easting       float64  // LKS-94 X coordinate
		northing      float64  // LKS-94 Y coordinate  
		expectedLat   float64  // Expected WGS-84 latitude
		expectedLon   float64  // Expected WGS-84 longitude
		tolerance     float64
	}{
		{
			name:        "Vilnius area",
			easting:     581234,
			northing:    6095678,
			expectedLat: 54.990387, // Lithuania latitude range: ~54-56°
			expectedLon: 25.269384, // Lithuania longitude range: ~21-27°
			tolerance:   0.0001,
		},
		{
			name:        "Kaunas area", 
			easting:     568123,
			northing:    6062456,
			expectedLat: 54.693908, // Lithuania latitude range
			expectedLon: 25.056723, // Lithuania longitude range
			tolerance:   0.0001,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lat, lon := LKS94ToWGS84(tc.easting, tc.northing)

			// Critical validation: coordinates should be in Lithuania, NOT Abu Dhabi
			if isInAbuDhabiArea(lat, lon) {
				t.Errorf("❌ COORDINATE MIXUP DETECTED: [%.6f, %.6f] appears to be in Abu Dhabi area!", lat, lon)
				t.Error("   This suggests lat/lon are swapped in the return values")
			}

			if !isInLithuania(lat, lon) {
				t.Errorf("❌ Coordinates [%.6f, %.6f] are not in Lithuania", lat, lon)
			}

			// Validate specific expected values
			if !isApproximatelyEqual(lat, tc.expectedLat, tc.tolerance) {
				t.Errorf("❌ Expected latitude %.6f, got %.6f (diff: %.6f)", 
					tc.expectedLat, lat, abs(lat-tc.expectedLat))
			}

			if !isApproximatelyEqual(lon, tc.expectedLon, tc.tolerance) {
				t.Errorf("❌ Expected longitude %.6f, got %.6f (diff: %.6f)", 
					tc.expectedLon, lon, abs(lon-tc.expectedLon))
			}

			t.Logf("✅ Transform correct: LKS-94 [%.0f, %.0f] -> WGS-84 [%.6f, %.6f]",
				tc.easting, tc.northing, lat, lon)
		})
	}
}

func TestCoordinateRangeValidation(t *testing.T) {
	// Test various Lithuanian coordinates to ensure they all stay in Lithuania
	lithuanianCoords := []struct {
		name     string
		easting  float64
		northing float64
	}{
		{"Vilnius center", 581234, 6095678},
		{"Kaunas center", 568123, 6062456}, 
		{"Klaipeda area", 317456, 6196543},
		{"Siauliai area", 486789, 6179234},
	}

	for _, coord := range lithuanianCoords {
		t.Run(coord.name, func(t *testing.T) {
			lat, lon := LKS94ToWGS84(coord.easting, coord.northing)

			if !isInLithuania(lat, lon) {
				t.Errorf("❌ %s coordinates [%.6f, %.6f] are outside Lithuania", 
					coord.name, lat, lon)
			}

			// Double-check not in Abu Dhabi (our previous bug)
			if isInAbuDhabiArea(lat, lon) {
				t.Errorf("❌ %s coordinates [%.6f, %.6f] are in Abu Dhabi - coordinate mixup!", 
					coord.name, lat, lon)
			}

			t.Logf("✅ %s: [%.6f, %.6f] is correctly in Lithuania", coord.name, lat, lon)
		})
	}
}

func TestReturnValueOrder(t *testing.T) {
	// Specific test to ensure function returns (latitude, longitude) as documented
	easting, northing := 581234.0, 6095678.0
	lat, lon := LKS94ToWGS84(easting, northing)

	// Lithuania is around 55°N, 25°E
	// If coordinates were swapped, we'd get ~25°N, 55°E (Abu Dhabi area)
	
	if lat < 25.0 {
		t.Errorf("❌ Latitude %.6f is too small - likely swapped with longitude", lat)
	}
	
	if lon > 50.0 {
		t.Errorf("❌ Longitude %.6f is too large - likely swapped with latitude", lon)
	}

	// Proper ranges for Lithuania
	if lat < 53.5 || lat > 56.5 {
		t.Errorf("❌ Latitude %.6f is outside Lithuania range (53.5-56.5)", lat)
	}
	
	if lon < 20.5 || lon > 27.0 {
		t.Errorf("❌ Longitude %.6f is outside Lithuania range (20.5-27.0)", lon)
	}

	t.Logf("✅ Return value order correct: latitude=%.6f, longitude=%.6f", lat, lon)
}

// Helper functions

func isInLithuania(lat, lon float64) bool {
	// Lithuania approximate boundaries
	return lat >= 53.5 && lat <= 56.5 && lon >= 20.5 && lon <= 27.0
}

func isInAbuDhabiArea(lat, lon float64) bool {
	// Abu Dhabi approximate area (where our bug put coordinates)
	return lat >= 24.0 && lat <= 25.0 && lon >= 54.0 && lon <= 56.0
}

func isApproximatelyEqual(a, b, tolerance float64) bool {
	return abs(a-b) <= tolerance
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}