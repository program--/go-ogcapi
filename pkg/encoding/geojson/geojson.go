package geojson

import "encoding/json"

type geojsonGeometry struct {
	Type        string           `json:"type"`
	Bounds      *json.RawMessage `json:"bbox,omitempty"`
	Coordinates *json.RawMessage `json:"coordinates,omitempty"`
	Geometries  *json.RawMessage `json:"geometries,omitempty"`
}

type geojsonFeature struct {
	Type       string           `json:"type"`
	Id         *string          `json:"id,omitempty"`
	Bounds     []float64        `json:"bbox,omitempty"`
	Geometry   *geojsonGeometry `json:"geometry"`
	Properties map[string]any   `json:"properties"`
}

type geojsonFeatureCollection struct {
	Type     string            `json:"type"`
	Bounds   []float64         `json:"bbox,omitempty"`
	Features []*geojsonFeature `json:"features"`
}
