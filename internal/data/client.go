// Package data provides HTTP clients for accessing Lithuanian road information APIs.
package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
)

// Client handles HTTP requests to Lithuanian traffic APIs
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new API client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		httpClient: httpClient,
	}
}

// FetchEALData fetches road restrictions from the EAL API
func (c *Client) FetchEALData() ([]EALLayer, error) {
	const url = "https://eismoinfo.lt/eismoinfo-backend/layer-dynamic-features/EAL?lks=true"

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch EAL data: %w", err)
	}
	defer resp.Body.Close()

	// Handle character encoding
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		reader = resp.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var layers []EALLayer
	if err := json.Unmarshal(body, &layers); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return layers, nil
}

// FetchArcGISData fetches speed control data from ArcGIS API
func (c *Client) FetchArcGISData() ([]ArcGISFeature, error) {
	// Get service information first
	maxRecords, err := c.getMaxRecordCount()
	if err != nil {
		return nil, fmt.Errorf("failed to get service info: %w", err)
	}

	// Fetch all features with pagination
	return c.fetchAllArcGISFeatures(maxRecords)
}

func (c *Client) getMaxRecordCount() (int, error) {
	const baseURL = "https://gis.ktvis.lt/arcgis/rest/services/PUB/PUB_ITS/MapServer/13"

	resp, err := c.httpClient.Get(baseURL + "?f=json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var info ArcGISServiceInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return 0, err
	}

	if info.MaxRecordCount > 0 {
		return info.MaxRecordCount, nil
	}
	return 1000, nil // default
}

func (c *Client) fetchAllArcGISFeatures(maxRecords int) ([]ArcGISFeature, error) {
	var allFeatures []ArcGISFeature
	offset := 0

	for {
		features, hasMore, err := c.fetchArcGISFeatureBatch(offset, maxRecords)
		if err != nil {
			return nil, err
		}

		allFeatures = append(allFeatures, features...)

		if !hasMore {
			break
		}

		offset += len(features)
	}

	return allFeatures, nil
}

func (c *Client) fetchArcGISFeatureBatch(offset, limit int) ([]ArcGISFeature, bool, error) {
	const queryURL = "https://gis.ktvis.lt/arcgis/rest/services/PUB/PUB_ITS/MapServer/13/query"

	// Build query parameters
	params := fmt.Sprintf("?where=1=1&outFields=*&returnGeometry=true&f=json&resultOffset=%d&resultRecordCount=%d&outSR=3346", offset, limit)

	resp, err := c.httpClient.Get(queryURL + params)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}

	var result ArcGISQueryResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, false, err
	}

	return result.Features, result.ExceededTransfer, nil
}
