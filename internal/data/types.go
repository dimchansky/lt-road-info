// Package data provides types and clients for accessing Lithuanian road information APIs.
package data

// EAL (Road Restrictions) Data Types

// EALLayer represents a layer from the EAL API response
type EALLayer struct {
	Layer    string       `json:"layer"`
	Name     string       `json:"name"`
	Features []EALFeature `json:"features"`
}

// EALFeature represents a road restriction feature
type EALFeature struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	Details      bool             `json:"details"`
	Icon         string           `json:"icon"`
	Points       []EALPoint       `json:"points"`
	Restrictions []EALRestriction `json:"restrictions"`
}

// EALPoint represents a point with min/max values
type EALPoint struct {
	Min   int       `json:"min"`
	Max   int       `json:"max"`
	Point []float64 `json:"point"`
}

// EALRestriction represents a specific restriction
type EALRestriction struct {
	ID        string   `json:"id"`
	Icon      string   `json:"icon"`
	IconValue float64  `json:"iconValue"`
	Lines     EALLines `json:"lines"`
}

// EALLines represents the geometry lines
type EALLines struct {
	Paths [][][]float64 `json:"paths"`
}

// ArcGIS (Speed Control) Data Types

// ArcGISFeature represents a speed control section from ArcGIS
type ArcGISFeature struct {
	Attributes map[string]interface{} `json:"attributes"`
	Geometry   ArcGISGeometry         `json:"geometry"`
}

// ArcGISGeometry represents ArcGIS geometry
type ArcGISGeometry struct {
	Paths [][][]float64 `json:"paths"`
}

// ArcGISQueryResponse represents the API response structure
type ArcGISQueryResponse struct {
	Features         []ArcGISFeature `json:"features"`
	ExceededTransfer bool            `json:"exceededTransferLimit"`
}

// ArcGISServiceInfo represents service metadata
type ArcGISServiceInfo struct {
	MaxRecordCount int `json:"maxRecordCount"`
}
