package geom_test

import (
	"reflect"
	"testing"

	"github.com/program--/go-ogcapi/pkg/geom"
)

func TestEmptyLineString(t *testing.T) {
	t.Parallel()

	empty_line := geom.NewLineString(geom.XY, nil)

	if empty_line.Type() != geom.LINESTRING {
		t.Errorf("LineString returned incorrect GeomType: %x\n", empty_line.Type())
	}

	if _, err := empty_line.PointAt(0); err == nil {
		t.Error("LineString with no points did not return an error from PointAt()\n")
	}

	if empty_line.NumPoints() != 0 {
		t.Error("LineString defined with nil data returned non-zero NumPoints()\n")
	}
}

func testHelperDimString(dim geom.DimsType) string {
	switch dim {
	case geom.XY:
		return "XY"
	case geom.XYZ:
		return "XYZ"
	case geom.XYM:
		return "XYM"
	case geom.XYZM:
		return "XYZM"
	default:
		return "UNKNOWN"
	}
}

func testHelperLineString(t *testing.T, dim geom.DimsType, numPoints int) {
	t.Parallel()

	// Construct data buffer
	line := geom.NewLineString(dim, nil)
	buffer := make([]float64, numPoints*line.Stride())

	val := 1.0
	for i := 0; i < numPoints; i++ {
		for j := 0; j < line.Stride(); j++ {
			buffer = append(buffer, val)
		}
		val += 1.0
	}

	// Begin tests
	line = geom.NewLineString(dim, buffer)
	dimStr := testHelperDimString(dim)
	if line.NumPoints() != numPoints {
		t.Errorf(
			"LineString with %d %s points returned incorrect count: %d\n",
			numPoints,
			dimStr,
			line.NumPoints(),
		)
	}

	for i := 0; i < (len(buffer) / line.Stride()); i++ {
		point, err := line.PointAt(i)
		if err != nil {
			t.Errorf("%s LineString returned error on point access: %s\n",
				dimStr,
				err.Error(),
			)
		}

		if point.Type() != geom.POINT {
			t.Errorf("Point %d from %s LineString has incorrect type: %x\n", i, dimStr, point.Type())
		}

		if point.Dimensions() != geom.XY {
			t.Errorf("Point %d from %s LineString is has incorrect dimensions: %x\n", i, dimStr, point.Dimensions())
		}

		if point.Empty() {
			t.Errorf("Point %d from %s LineString is empty\n", i, dimStr)
		}

		coord := float64(i) + 1.0
		if point.X() != coord || point.Y() != coord {
			t.Errorf(
				"Point %d from %s LineString has incorrect coords: (%f, %f) != (%f, %f)",
				i, dimStr,
				coord, coord,
				point.X(), point.Y(),
			)
		}

		ptr := &buffer[i*line.Stride()]
		if !reflect.DeepEqual(ptr, &point.Coords()[0]) {
			t.Errorf(
				"Point %d and %s LineString do not share same backing buffer: %p != %p",
				i, dimStr,
				ptr, &point.Coords()[0],
			)
		}
	}

}

func TestXYLineString(t *testing.T) {
	t.Parallel()

	buffer := []float64{1.0, 1.0, 2.0, 2.0, 3.0, 3.0, 4.0, 4.0}
	line := geom.NewLineString(geom.XY, buffer)

	if line.NumPoints() != 4 {
		t.Errorf("LineString with 4 XY points returned incorrect count: %d\n", line.NumPoints())
	}

	for i := 0; i < (len(buffer) / line.Stride()); i++ {
		point, err := line.PointAt(i)
		if err != nil {
			t.Errorf("LineString with 4 XY points returned error on access: %s\n", err.Error())
		}

		if point.Type() != geom.POINT {
			t.Errorf("Point %d from XY LineString has incorrect type: %x\n", i, point.Type())
		}

		if point.Dimensions() != geom.XY {
			t.Errorf("Point %d from XY LineString is has incorrect dimensions: %x\n", i, point.Dimensions())
		}

		if point.Empty() {
			t.Errorf("Point %d from XY LineString is empty\n", i)
		}

		coord := float64(i) + 1.0
		if point.X() != coord || point.Y() != coord {
			t.Errorf(
				"Point %d from XY LineString has incorrect coords: (%f, %f) != (%f, %f)",
				i,
				coord, coord,
				point.X(), point.Y(),
			)
		}

		ptr := &buffer[i*line.Stride()]
		if !reflect.DeepEqual(ptr, &point.Coords()[0]) {
			t.Errorf("Point %d and XY LineString do not share same backing buffer: %p != %p", i, ptr, &point.Coords()[0])
		}
	}
}

func TestXYZLineString(t *testing.T) {
	t.Parallel()

	buffer := []float64{1.0, 1.0, 1.0, 2.0, 2.0, 2.0, 3.0, 3.0, 3.0, 4.0, 4.0, 4.0}
	line := geom.NewLineString(geom.XYZ, buffer)

	if line.NumPoints() != 4 {
		t.Errorf("LineString with 4 XYZ points returned incorrect count: %d\n", line.NumPoints())
	}

	for i := 0; i < (len(buffer) / line.Stride()); i++ {
		point, err := line.PointAt(i)
		if err != nil {
			t.Errorf("LineString with 4 XYZ points returned error on access: %s\n", err.Error())
		}

		if point.Type() != geom.POINT {
			t.Errorf("Point %d from XYZ LineString has incorrect type: %x\n", i, point.Type())
		}

		if point.Dimensions() != geom.XYZ {
			t.Errorf("Point %d from XYZ LineString is has incorrect dimensions: %x\n", i, point.Dimensions())
		}

		if point.Empty() {
			t.Errorf("Point %d from XYZ LineString is empty\n", i)
		}

		coord := float64(i) + 1.0
		if point.X() != coord || point.Y() != coord || point.Z() != coord {
			t.Errorf(
				"Point %d from XYZ LineString has incorrect coords: (%f, %f, %f) != (%f, %f, %f)",
				i,
				coord, coord, coord,
				point.X(), point.Y(), point.Z(),
			)
		}

		ptr := &buffer[i*line.Stride()]
		if !reflect.DeepEqual(ptr, &point.Coords()[0]) {
			t.Errorf("Point %d and XYZ LineString do not share same backing buffer: %p != %p", i, ptr, &point.Coords()[0])
		}
	}
}
