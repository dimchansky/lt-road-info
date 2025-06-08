package transform

import (
	"github.com/wroge/wgs84/v2"
)

// Cached transformation function from LKS-94 (EPSG:3346) to WGS84 (EPSG:4326)
var cachedTransform wgs84.Func

func init() {
	// Define the geocentric CRS (EPSG:4978)
	epsg4978 := base{}

	// Define the geographic CRS for ETRS89 (EPSG:4258)
	epsg4258 := wgs84.Geographic(epsg4978, wgs84.NewSpheroid(6378137, 298.257222101))

	// Define the projected CRS for LKS-94 (EPSG:3346)
	epsg3346 := wgs84.TransverseMercator(epsg4258, 24, 0, 0.9998, 500000, 0)

	// Define the geographic CRS for WGS84 (EPSG:4326)
	epsg4326 := wgs84.Geographic(epsg4978, wgs84.NewSpheroid(6378137, 298.257223563))

	// Initialize the cached transformation function
	cachedTransform = wgs84.Transform(epsg3346, epsg4326)
}

// LKS94ToWGS84 transforms LKS-94 (EPSG:3346) coordinates to WGS84 (EPSG:4326).
func LKS94ToWGS84(easting, northing float64) (latitude, longitude float64) {
	// Use the cached transformation function
	longitude, latitude, _ = cachedTransform(easting, northing, 0.0)
	return latitude, longitude
}

// base struct implements the wgs84.CRS interface for geocentric CRS
type base struct{}

func (base) Base() wgs84.CRS {
	return nil
}

func (base) Spheroid() wgs84.Spheroid {
	return wgs84.Spheroid{}
}

func (base) ToBase(x0, y0, z0 float64) (float64, float64, float64) {
	return x0, y0, z0
}

func (base) FromBase(x0, y0, z0 float64) (float64, float64, float64) {
	return x0, y0, z0
}