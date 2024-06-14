//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/oapi-codegen/runtime"
)

// Defines values for ExtentSpatialCrs.
const (
	HttpwwwOpengisNetdefcrsOGC13CRS84 ExtentSpatialCrs = "http://www.opengis.net/def/crs/OGC/1.3/CRS84"
)

// Defines values for ExtentTemporalTrs.
const (
	HttpwwwOpengisNetdefuomISO86010Gregorian ExtentTemporalTrs = "http://www.opengis.net/def/uom/ISO-8601/0/Gregorian"
)

// Defines values for FeatureCollectionGeoJSONType.
const (
	FeatureCollection FeatureCollectionGeoJSONType = "FeatureCollection"
)

// Defines values for FeatureGeoJSONType.
const (
	FeatureGeoJSONTypeFeature FeatureGeoJSONType = "Feature"
)

// Defines values for GeometrycollectionGeoJSONType.
const (
	GeometryCollection GeometrycollectionGeoJSONType = "GeometryCollection"
)

// Defines values for LinestringGeoJSONType.
const (
	LineString LinestringGeoJSONType = "LineString"
)

// Defines values for MultilinestringGeoJSONType.
const (
	MultiLineString MultilinestringGeoJSONType = "MultiLineString"
)

// Defines values for MultipointGeoJSONType.
const (
	MultiPoint MultipointGeoJSONType = "MultiPoint"
)

// Defines values for MultipolygonGeoJSONType.
const (
	MultiPolygon MultipolygonGeoJSONType = "MultiPolygon"
)

// Defines values for PointGeoJSONType.
const (
	Point PointGeoJSONType = "Point"
)

// Defines values for PolygonGeoJSONType.
const (
	Polygon PolygonGeoJSONType = "Polygon"
)

// Collection defines model for collection.
type Collection struct {
	// Crs the list of coordinate reference systems supported by the service
	Crs *[]string `json:"crs,omitempty"`

	// Description a description of the features in the collection
	Description *string `json:"description,omitempty"`

	// Extent The extent of the features in the collection. In the Core only spatial and temporal
	// extents are specified. Extensions may add additional members to represent other
	// extents, for example, thermal or pressure ranges.
	Extent *Extent `json:"extent,omitempty"`

	// Id identifier of the collection used, for example, in URIs
	Id string `json:"id"`

	// ItemType indicator about the type of the items in the collection (the default value is 'feature').
	ItemType      *string         `json:"itemType,omitempty"`
	LinkTemplates *[]LinkTemplate `json:"linkTemplates,omitempty"`
	Links         []Link          `json:"links"`

	// Title human readable title of the collection
	Title *string `json:"title,omitempty"`
}

// Collections defines model for collections.
type Collections struct {
	Collections []Collection `json:"collections"`
	Links       []Link       `json:"links"`
}

// ConfClasses defines model for confClasses.
type ConfClasses struct {
	ConformsTo []string `json:"conformsTo"`
}

// Exception Information about the exception: an error code plus an optional description.
type Exception struct {
	Code        string  `json:"code"`
	Description *string `json:"description,omitempty"`
}

// Extent The extent of the features in the collection. In the Core only spatial and temporal
// extents are specified. Extensions may add additional members to represent other
// extents, for example, thermal or pressure ranges.
type Extent struct {
	// Spatial The spatial extent of the features in the collection.
	Spatial *struct {
		// Bbox One or more bounding boxes that describe the spatial extent of the dataset.
		// In the Core only a single bounding box is supported. Extensions may support
		// additional areas. If multiple areas are provided, the union of the bounding
		// boxes describes the spatial extent.
		Bbox *[][]float32 `json:"bbox,omitempty"`

		// Crs Coordinate reference system of the coordinates in the spatial extent
		// (property `bbox`). The default reference system is WGS 84 longitude/latitude.
		// In the Core this is the only supported coordinate reference system.
		// Extensions may support additional coordinate reference systems and add
		// additional enum values.
		Crs *ExtentSpatialCrs `json:"crs,omitempty"`
	} `json:"spatial,omitempty"`

	// Temporal The temporal extent of the features in the collection.
	Temporal *struct {
		// Interval One or more time intervals that describe the temporal extent of the dataset.
		// The value `null` is supported and indicates an unbounded interval end.
		// In the Core only a single time interval is supported. Extensions may support
		// multiple intervals. If multiple intervals are provided, the union of the
		// intervals describes the temporal extent.
		Interval *[][]time.Time `json:"interval,omitempty"`

		// Trs Coordinate reference system of the coordinates in the temporal extent
		// (property `interval`). The default reference system is the Gregorian calendar.
		// In the Core this is the only supported temporal coordinate reference system.
		// Extensions may support additional temporal coordinate reference systems and add
		// additional enum values.
		Trs *ExtentTemporalTrs `json:"trs,omitempty"`
	} `json:"temporal,omitempty"`
}

// ExtentSpatialCrs Coordinate reference system of the coordinates in the spatial extent
// (property `bbox`). The default reference system is WGS 84 longitude/latitude.
// In the Core this is the only supported coordinate reference system.
// Extensions may support additional coordinate reference systems and add
// additional enum values.
type ExtentSpatialCrs string

// ExtentTemporalTrs Coordinate reference system of the coordinates in the temporal extent
// (property `interval`). The default reference system is the Gregorian calendar.
// In the Core this is the only supported temporal coordinate reference system.
// Extensions may support additional temporal coordinate reference systems and add
// additional enum values.
type ExtentTemporalTrs string

// FeatureCollectionGeoJSON defines model for featureCollectionGeoJSON.
type FeatureCollectionGeoJSON struct {
	Features []FeatureGeoJSON `json:"features"`
	Links    *[]Link          `json:"links,omitempty"`

	// NumberMatched The number of features of the feature type that match the selection
	// parameters like `bbox`.
	NumberMatched *NumberMatched `json:"numberMatched,omitempty"`

	// NumberReturned The number of features in the feature collection.
	//
	// A server may omit this information in a response, if the information
	// about the number of features is not known or difficult to compute.
	//
	// If the value is provided, the value shall be identical to the number
	// of items in the "features" array.
	NumberReturned *NumberReturned `json:"numberReturned,omitempty"`

	// TimeStamp This property indicates the time and date when the response was generated.
	TimeStamp *TimeStamp                   `json:"timeStamp,omitempty"`
	Type      FeatureCollectionGeoJSONType `json:"type"`
}

// FeatureCollectionGeoJSONType defines model for FeatureCollectionGeoJSON.Type.
type FeatureCollectionGeoJSONType string

// FeatureGeoJSON defines model for featureGeoJSON.
type FeatureGeoJSON struct {
	Geometry   GeometryGeoJSON         `json:"geometry"`
	Id         *FeatureGeoJSON_Id      `json:"id,omitempty"`
	Links      *[]Link                 `json:"links,omitempty"`
	Properties *map[string]interface{} `json:"properties"`
	Type       FeatureGeoJSONType      `json:"type"`
}

// FeatureGeoJSONId0 defines model for .
type FeatureGeoJSONId0 = string

// FeatureGeoJSONId1 defines model for .
type FeatureGeoJSONId1 = int

// FeatureGeoJSON_Id defines model for FeatureGeoJSON.Id.
type FeatureGeoJSON_Id struct {
	union json.RawMessage
}

// FeatureGeoJSONType defines model for FeatureGeoJSON.Type.
type FeatureGeoJSONType string

// GeometryGeoJSON defines model for geometryGeoJSON.
type GeometryGeoJSON struct {
	union json.RawMessage
}

// GeometrycollectionGeoJSON defines model for geometrycollectionGeoJSON.
type GeometrycollectionGeoJSON struct {
	Geometries []GeometryGeoJSON             `json:"geometries"`
	Type       GeometrycollectionGeoJSONType `json:"type"`
}

// GeometrycollectionGeoJSONType defines model for GeometrycollectionGeoJSON.Type.
type GeometrycollectionGeoJSONType string

// LandingPage defines model for landingPage.
type LandingPage struct {
	Description *string `json:"description,omitempty"`
	Links       []Link  `json:"links"`
	Title       *string `json:"title,omitempty"`
}

// LinestringGeoJSON defines model for linestringGeoJSON.
type LinestringGeoJSON struct {
	Coordinates [][]float32           `json:"coordinates"`
	Type        LinestringGeoJSONType `json:"type"`
}

// LinestringGeoJSONType defines model for LinestringGeoJSON.Type.
type LinestringGeoJSONType string

// Link defines model for link.
type Link struct {
	Href     string  `json:"href"`
	Hreflang *string `json:"hreflang,omitempty"`
	Length   *int    `json:"length,omitempty"`
	Rel      *string `json:"rel,omitempty"`
	Title    *string `json:"title,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// LinkTemplate defines model for linkTemplate.
type LinkTemplate struct {
	Rel         string  `json:"rel"`
	Title       *string `json:"title,omitempty"`
	Type        *string `json:"type,omitempty"`
	UriTemplate string  `json:"uriTemplate"`
	VarBase     *string `json:"varBase,omitempty"`
}

// MultilinestringGeoJSON defines model for multilinestringGeoJSON.
type MultilinestringGeoJSON struct {
	Coordinates [][][]float32              `json:"coordinates"`
	Type        MultilinestringGeoJSONType `json:"type"`
}

// MultilinestringGeoJSONType defines model for MultilinestringGeoJSON.Type.
type MultilinestringGeoJSONType string

// MultipointGeoJSON defines model for multipointGeoJSON.
type MultipointGeoJSON struct {
	Coordinates [][]float32           `json:"coordinates"`
	Type        MultipointGeoJSONType `json:"type"`
}

// MultipointGeoJSONType defines model for MultipointGeoJSON.Type.
type MultipointGeoJSONType string

// MultipolygonGeoJSON defines model for multipolygonGeoJSON.
type MultipolygonGeoJSON struct {
	Coordinates [][][][]float32         `json:"coordinates"`
	Type        MultipolygonGeoJSONType `json:"type"`
}

// MultipolygonGeoJSONType defines model for MultipolygonGeoJSON.Type.
type MultipolygonGeoJSONType string

// NumberMatched The number of features of the feature type that match the selection
// parameters like `bbox`.
type NumberMatched = int

// NumberReturned The number of features in the feature collection.
//
// A server may omit this information in a response, if the information
// about the number of features is not known or difficult to compute.
//
// If the value is provided, the value shall be identical to the number
// of items in the "features" array.
type NumberReturned = int

// PointGeoJSON defines model for pointGeoJSON.
type PointGeoJSON struct {
	Coordinates []float32        `json:"coordinates"`
	Type        PointGeoJSONType `json:"type"`
}

// PointGeoJSONType defines model for PointGeoJSON.Type.
type PointGeoJSONType string

// PolygonGeoJSON defines model for polygonGeoJSON.
type PolygonGeoJSON struct {
	Coordinates [][][]float32      `json:"coordinates"`
	Type        PolygonGeoJSONType `json:"type"`
}

// PolygonGeoJSONType defines model for PolygonGeoJSON.Type.
type PolygonGeoJSONType string

// TimeStamp This property indicates the time and date when the response was generated.
type TimeStamp = time.Time

// Bbox defines model for bbox.
type Bbox = []float32

// CollectionId defines model for collectionId.
type CollectionId = string

// Datetime defines model for datetime.
type Datetime = string

// FeatureId defines model for featureId.
type FeatureId = string

// Limit defines model for limit.
type Limit = int

// ConformanceDeclaration defines model for ConformanceDeclaration.
type ConformanceDeclaration = ConfClasses

// Feature defines model for Feature.
type Feature = FeatureGeoJSON

// Features defines model for Features.
type Features = FeatureCollectionGeoJSON

// InvalidParameter Information about the exception: an error code plus an optional description.
type InvalidParameter = Exception

// ServerError Information about the exception: an error code plus an optional description.
type ServerError = Exception

// GetFeaturesParams defines parameters for GetFeatures.
type GetFeaturesParams struct {
	// Limit The optional limit parameter limits the number of items that are presented in the response document.
	//
	// Only items are counted that are on the first level of the collection in the response document.
	// Nested objects contained within the explicitly requested items shall not be counted.
	//
	// Minimum = 1. Maximum = 10000. Default = 10.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Bbox Only features that have a geometry that intersects the bounding box are selected.
	// The bounding box is provided as four or six numbers, depending on whether the
	// coordinate reference system includes a vertical axis (height or depth):
	//
	// * Lower left corner, coordinate axis 1
	// * Lower left corner, coordinate axis 2
	// * Minimum value, coordinate axis 3 (optional)
	// * Upper right corner, coordinate axis 1
	// * Upper right corner, coordinate axis 2
	// * Maximum value, coordinate axis 3 (optional)
	//
	// If the value consists of four numbers, the coordinate reference system is
	// WGS 84 longitude/latitude (http://www.opengis.net/def/crs/OGC/1.3/CRS84)
	// unless a different coordinate reference system is specified in the parameter `bbox-crs`.
	//
	// If the value consists of six numbers, the coordinate reference system is WGS 84
	// longitude/latitude/ellipsoidal height (http://www.opengis.net/def/crs/OGC/0/CRS84h)
	// unless a different coordinate reference system is specified in the parameter `bbox-crs`.
	//
	// The query parameter `bbox-crs` is specified in OGC API - Features - Part 2: Coordinate
	// Reference Systems by Reference.
	//
	// For WGS 84 longitude/latitude the values are in most cases the sequence of
	// minimum longitude, minimum latitude, maximum longitude and maximum latitude.
	// However, in cases where the box spans the antimeridian the first value
	// (west-most box edge) is larger than the third or fourth value (east-most box edge).
	//
	// If the vertical axis is included, the third and the sixth number are
	// the bottom and the top of the 3-dimensional bounding box.
	//
	// If a feature has multiple spatial geometry properties, it is the decision of the
	// server whether only a single spatial geometry property is used to determine
	// the extent or all relevant geometries.
	Bbox *Bbox `form:"bbox,omitempty" json:"bbox,omitempty"`

	// Datetime Either a date-time or an interval. Date and time expressions adhere to RFC 3339.
	// Intervals may be bounded or half-bounded (double-dots at start or end).
	//
	// Examples:
	//
	// * A date-time: "2018-02-12T23:20:50Z"
	// * A bounded interval: "2018-02-12T00:00:00Z/2018-03-18T12:31:12Z"
	// * Half-bounded intervals: "2018-02-12T00:00:00Z/.." or "../2018-03-18T12:31:12Z"
	//
	// Only features that have a temporal property that intersects the value of
	// `datetime` are selected.
	//
	// If a feature has multiple temporal properties, it is the decision of the
	// server whether only a single temporal property is used to determine
	// the extent or all relevant temporal properties.
	Datetime *Datetime `form:"datetime,omitempty" json:"datetime,omitempty"`
}

// AsFeatureGeoJSONId0 returns the union data inside the FeatureGeoJSON_Id as a FeatureGeoJSONId0
func (t FeatureGeoJSON_Id) AsFeatureGeoJSONId0() (FeatureGeoJSONId0, error) {
	var body FeatureGeoJSONId0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFeatureGeoJSONId0 overwrites any union data inside the FeatureGeoJSON_Id as the provided FeatureGeoJSONId0
func (t *FeatureGeoJSON_Id) FromFeatureGeoJSONId0(v FeatureGeoJSONId0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFeatureGeoJSONId0 performs a merge with any union data inside the FeatureGeoJSON_Id, using the provided FeatureGeoJSONId0
func (t *FeatureGeoJSON_Id) MergeFeatureGeoJSONId0(v FeatureGeoJSONId0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFeatureGeoJSONId1 returns the union data inside the FeatureGeoJSON_Id as a FeatureGeoJSONId1
func (t FeatureGeoJSON_Id) AsFeatureGeoJSONId1() (FeatureGeoJSONId1, error) {
	var body FeatureGeoJSONId1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFeatureGeoJSONId1 overwrites any union data inside the FeatureGeoJSON_Id as the provided FeatureGeoJSONId1
func (t *FeatureGeoJSON_Id) FromFeatureGeoJSONId1(v FeatureGeoJSONId1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFeatureGeoJSONId1 performs a merge with any union data inside the FeatureGeoJSON_Id, using the provided FeatureGeoJSONId1
func (t *FeatureGeoJSON_Id) MergeFeatureGeoJSONId1(v FeatureGeoJSONId1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FeatureGeoJSON_Id) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FeatureGeoJSON_Id) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPointGeoJSON returns the union data inside the GeometryGeoJSON as a PointGeoJSON
func (t GeometryGeoJSON) AsPointGeoJSON() (PointGeoJSON, error) {
	var body PointGeoJSON
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPointGeoJSON overwrites any union data inside the GeometryGeoJSON as the provided PointGeoJSON
func (t *GeometryGeoJSON) FromPointGeoJSON(v PointGeoJSON) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePointGeoJSON performs a merge with any union data inside the GeometryGeoJSON, using the provided PointGeoJSON
func (t *GeometryGeoJSON) MergePointGeoJSON(v PointGeoJSON) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultipointGeoJSON returns the union data inside the GeometryGeoJSON as a MultipointGeoJSON
func (t GeometryGeoJSON) AsMultipointGeoJSON() (MultipointGeoJSON, error) {
	var body MultipointGeoJSON
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultipointGeoJSON overwrites any union data inside the GeometryGeoJSON as the provided MultipointGeoJSON
func (t *GeometryGeoJSON) FromMultipointGeoJSON(v MultipointGeoJSON) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultipointGeoJSON performs a merge with any union data inside the GeometryGeoJSON, using the provided MultipointGeoJSON
func (t *GeometryGeoJSON) MergeMultipointGeoJSON(v MultipointGeoJSON) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsLinestringGeoJSON returns the union data inside the GeometryGeoJSON as a LinestringGeoJSON
func (t GeometryGeoJSON) AsLinestringGeoJSON() (LinestringGeoJSON, error) {
	var body LinestringGeoJSON
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLinestringGeoJSON overwrites any union data inside the GeometryGeoJSON as the provided LinestringGeoJSON
func (t *GeometryGeoJSON) FromLinestringGeoJSON(v LinestringGeoJSON) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLinestringGeoJSON performs a merge with any union data inside the GeometryGeoJSON, using the provided LinestringGeoJSON
func (t *GeometryGeoJSON) MergeLinestringGeoJSON(v LinestringGeoJSON) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultilinestringGeoJSON returns the union data inside the GeometryGeoJSON as a MultilinestringGeoJSON
func (t GeometryGeoJSON) AsMultilinestringGeoJSON() (MultilinestringGeoJSON, error) {
	var body MultilinestringGeoJSON
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultilinestringGeoJSON overwrites any union data inside the GeometryGeoJSON as the provided MultilinestringGeoJSON
func (t *GeometryGeoJSON) FromMultilinestringGeoJSON(v MultilinestringGeoJSON) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultilinestringGeoJSON performs a merge with any union data inside the GeometryGeoJSON, using the provided MultilinestringGeoJSON
func (t *GeometryGeoJSON) MergeMultilinestringGeoJSON(v MultilinestringGeoJSON) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsPolygonGeoJSON returns the union data inside the GeometryGeoJSON as a PolygonGeoJSON
func (t GeometryGeoJSON) AsPolygonGeoJSON() (PolygonGeoJSON, error) {
	var body PolygonGeoJSON
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPolygonGeoJSON overwrites any union data inside the GeometryGeoJSON as the provided PolygonGeoJSON
func (t *GeometryGeoJSON) FromPolygonGeoJSON(v PolygonGeoJSON) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePolygonGeoJSON performs a merge with any union data inside the GeometryGeoJSON, using the provided PolygonGeoJSON
func (t *GeometryGeoJSON) MergePolygonGeoJSON(v PolygonGeoJSON) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMultipolygonGeoJSON returns the union data inside the GeometryGeoJSON as a MultipolygonGeoJSON
func (t GeometryGeoJSON) AsMultipolygonGeoJSON() (MultipolygonGeoJSON, error) {
	var body MultipolygonGeoJSON
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMultipolygonGeoJSON overwrites any union data inside the GeometryGeoJSON as the provided MultipolygonGeoJSON
func (t *GeometryGeoJSON) FromMultipolygonGeoJSON(v MultipolygonGeoJSON) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMultipolygonGeoJSON performs a merge with any union data inside the GeometryGeoJSON, using the provided MultipolygonGeoJSON
func (t *GeometryGeoJSON) MergeMultipolygonGeoJSON(v MultipolygonGeoJSON) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsGeometrycollectionGeoJSON returns the union data inside the GeometryGeoJSON as a GeometrycollectionGeoJSON
func (t GeometryGeoJSON) AsGeometrycollectionGeoJSON() (GeometrycollectionGeoJSON, error) {
	var body GeometrycollectionGeoJSON
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGeometrycollectionGeoJSON overwrites any union data inside the GeometryGeoJSON as the provided GeometrycollectionGeoJSON
func (t *GeometryGeoJSON) FromGeometrycollectionGeoJSON(v GeometrycollectionGeoJSON) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGeometrycollectionGeoJSON performs a merge with any union data inside the GeometryGeoJSON, using the provided GeometrycollectionGeoJSON
func (t *GeometryGeoJSON) MergeGeometrycollectionGeoJSON(v GeometrycollectionGeoJSON) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t GeometryGeoJSON) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GeometryGeoJSON) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// landing page
	// (GET /)
	GetLandingPage(w http.ResponseWriter, r *http.Request)
	// the feature collections in the dataset
	// (GET /collections)
	GetCollections(w http.ResponseWriter, r *http.Request)
	// describe the feature collection with id `collectionId`
	// (GET /collections/{collectionId})
	DescribeCollection(w http.ResponseWriter, r *http.Request, collectionId CollectionId)
	// fetch features
	// (GET /collections/{collectionId}/items)
	GetFeatures(w http.ResponseWriter, r *http.Request, collectionId CollectionId, params GetFeaturesParams)
	// fetch a single feature
	// (GET /collections/{collectionId}/items/{featureId})
	GetFeature(w http.ResponseWriter, r *http.Request, collectionId CollectionId, featureId FeatureId)
	// information about specifications that this API conforms to
	// (GET /conformance)
	GetConformanceDeclaration(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetLandingPage operation middleware
func (siw *ServerInterfaceWrapper) GetLandingPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetLandingPage(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetCollections operation middleware
func (siw *ServerInterfaceWrapper) GetCollections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCollections(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DescribeCollection operation middleware
func (siw *ServerInterfaceWrapper) DescribeCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "collectionId" -------------
	var collectionId CollectionId

	err = runtime.BindStyledParameterWithOptions("simple", "collectionId", r.PathValue("collectionId"), &collectionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "collectionId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DescribeCollection(w, r, collectionId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetFeatures operation middleware
func (siw *ServerInterfaceWrapper) GetFeatures(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "collectionId" -------------
	var collectionId CollectionId

	err = runtime.BindStyledParameterWithOptions("simple", "collectionId", r.PathValue("collectionId"), &collectionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "collectionId", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFeaturesParams

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", false, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "bbox" -------------

	err = runtime.BindQueryParameter("form", false, false, "bbox", r.URL.Query(), &params.Bbox)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "bbox", Err: err})
		return
	}

	// ------------- Optional query parameter "datetime" -------------

	err = runtime.BindQueryParameter("form", false, false, "datetime", r.URL.Query(), &params.Datetime)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "datetime", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFeatures(w, r, collectionId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetFeature operation middleware
func (siw *ServerInterfaceWrapper) GetFeature(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "collectionId" -------------
	var collectionId CollectionId

	err = runtime.BindStyledParameterWithOptions("simple", "collectionId", r.PathValue("collectionId"), &collectionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "collectionId", Err: err})
		return
	}

	// ------------- Path parameter "featureId" -------------
	var featureId FeatureId

	err = runtime.BindStyledParameterWithOptions("simple", "featureId", r.PathValue("featureId"), &featureId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "featureId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFeature(w, r, collectionId, featureId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetConformanceDeclaration operation middleware
func (siw *ServerInterfaceWrapper) GetConformanceDeclaration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetConformanceDeclaration(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       *http.ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m *http.ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m *http.ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/", wrapper.GetLandingPage)
	m.HandleFunc("GET "+options.BaseURL+"/collections", wrapper.GetCollections)
	m.HandleFunc("GET "+options.BaseURL+"/collections/{collectionId}", wrapper.DescribeCollection)
	m.HandleFunc("GET "+options.BaseURL+"/collections/{collectionId}/items", wrapper.GetFeatures)
	m.HandleFunc("GET "+options.BaseURL+"/collections/{collectionId}/items/{featureId}", wrapper.GetFeature)
	m.HandleFunc("GET "+options.BaseURL+"/conformance", wrapper.GetConformanceDeclaration)

	return m
}
