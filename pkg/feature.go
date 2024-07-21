package ogc_api

import (
	"fmt"
	"reflect"

	"github.com/program--/go-ogcapi/pkg/geom"
)

type PropertyType interface {
	~int | ~float64 | ~bool | ~string
}

func getProperty[T PropertyType](properties map[string]any, key string) (T, error) {
	generic_property, ok := properties[key]
	if !ok {
		return *new(T), fmt.Errorf("key '%s' does not exist in properties", key)
	}

	property, ok := generic_property.(T)
	if !ok {
		return *new(T), fmt.Errorf("property '%s' is not coercible to type of `%s`", key, reflect.TypeOf(property).String())
	}

	return property, nil
}

type Feature struct {
	Geometry   geom.Geometry
	properties map[string]any
}

type FeatureCollection = []*Feature

func (feature *Feature) PropertyBool(key string) (bool, error) {
	return getProperty[bool](feature.properties, key)
}

func (feature *Feature) PropertyFloat64(key string) (float64, error) {
	return getProperty[float64](feature.properties, key)
}

func (feature *Feature) PropertyInt(key string) (int, error) {
	return getProperty[int](feature.properties, key)
}

func (feature *Feature) PropertyString(key string) (string, error) {
	return getProperty[string](feature.properties, key)
}

func (feature *Feature) SetProperty(key string, value any) {
	feature.properties[key] = value
}
