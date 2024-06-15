package geom_test

import (
	"fmt"
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

func TestXYLineString(t *testing.T) {
	t.Parallel()
	testHelperLineString(t, geom.XY, 10)
}

func TestXYZLineString(t *testing.T) {
	t.Parallel()
	testHelperLineString(t, geom.XYZ, 10)
}

func TestXYMLineString(t *testing.T) {
	t.Parallel()
	testHelperLineString(t, geom.XYM, 10)
}

func TestXYZMLineString(t *testing.T) {
	t.Parallel()
	testHelperLineString(t, geom.XYZM, 10)
}

// ============================================================================
// Test Helpers ===============================================================
// ============================================================================
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
	// Construct data buffer
	line := geom.NewLineString(dim, nil)
	buffer := make([]float64, 0, numPoints*line.Stride())

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

	for i := 0; i < numPoints; i++ {
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

		if point.Dimensions() != dim {
			t.Errorf(
				"Point %d from %s LineString is has incorrect dimensions: %s\n",
				i, dimStr,
				testHelperDimString(line.Dimensions()),
			)
		}

		if point.Empty() {
			t.Errorf("Point %d from %s LineString is empty\n", i, dimStr)
		}

		var (
			coordStr, pointStr string
			coordEqual         bool
		)
		coord := float64(i) + 1.0
		switch dim {
		case geom.XY:
			coordStr = fmt.Sprintf("%f, %f", coord, coord)
			pointStr = fmt.Sprintf("%f, %f", point.X(), point.Y())
			coordEqual = point.X() == coord && point.Y() == coord
		case geom.XYZ:
			coordStr = fmt.Sprintf("%f, %f, %f", coord, coord, coord)
			pointStr = fmt.Sprintf("%f, %f, %f", point.X(), point.Y(), point.Z())
			coordEqual = point.X() == coord && point.Y() == coord && point.Z() == coord
		case geom.XYM:
			coordStr = fmt.Sprintf("%f, %f, %f", coord, coord, coord)
			pointStr = fmt.Sprintf("%f, %f, %f", point.X(), point.Y(), point.M())
			coordEqual = point.X() == coord && point.Y() == coord && point.M() == coord
		case geom.XYZM:
			coordStr = fmt.Sprintf("%f, %f, %f, %f", coord, coord, coord, coord)
			pointStr = fmt.Sprintf("%f, %f, %f, %f", point.X(), point.Y(), point.Z(), point.M())
			coordEqual = point.X() == coord && point.Y() == coord && point.Z() == coord && point.M() == coord
		}

		if !coordEqual {
			t.Errorf(
				"Point %d from %s LineString has incorrect coords: %s != %s",
				i, dimStr,
				coordStr,
				pointStr,
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
