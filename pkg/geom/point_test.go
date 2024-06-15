package geom_test

import (
	"math"
	"testing"

	"github.com/program--/go-ogcapi/pkg/geom"
)

func TestEmptyPoint(t *testing.T) {
	t.Parallel()

	empty_point := geom.NewPoint(geom.XY, []float64{})
	nil_point := geom.NewPoint(geom.XY, nil)

	if empty_point.Type() != geom.POINT {
		t.Errorf("Point returned incorrect GeomType: %x\n", empty_point.Type())
	}

	if !nil_point.Empty() {
		t.Error("Point with data == nil returned !Empty()\n")
	}

	if !empty_point.Empty() {
		t.Error("Point with len(data) == 0 returned !Empty()\n")
	}

	if empty_point.Dimensions() != geom.XY {
		t.Errorf("Point with dim XY (0) returned %d", empty_point.Dimensions())
	}
}

func TestXYPoint(t *testing.T) {
	t.Parallel()

	point := geom.NewPoint(geom.XY, []float64{1.0, 2.0})

	if point.Empty() {
		t.Error("Point with data {1.0, 2.0} returned Empty()\n")
	}

	if point.X() != 1.0 {
		t.Errorf("Point returned incorrect X: %f\n", point.X())
	}

	if point.Y() != 2.0 {
		t.Errorf("Point returned incorrect Y: %f\n", point.Y())
	}

	if !math.IsNaN(point.Z()) {
		t.Errorf("Point returned incorrect Z: %f\n", point.Z())
	}

	if !math.IsNaN(point.M()) {
		t.Errorf("Point returned incorrect M: %f\n", point.M())
	}

	if point.Stride() != 2 {
		t.Errorf("Point with dim XY returned stride %d instead of 2", point.Stride())
	}
}

func TestXYZPoint(t *testing.T) {
	t.Parallel()

	point := geom.NewPoint(geom.XYZ, []float64{1.0, 2.0, 3.0})

	if point.Empty() {
		t.Error("Point with data {1.0, 2.0, 3.0} returned Empty()\n")
	}

	if point.X() != 1.0 {
		t.Errorf("Point returned incorrect X: %f\n", point.X())
	}

	if point.Y() != 2.0 {
		t.Errorf("Point returned incorrect Y: %f\n", point.Y())
	}

	if point.Z() != 3.0 {
		t.Errorf("Point returned incorrect Z: %f\n", point.Z())
	}

	if !math.IsNaN(point.M()) {
		t.Errorf("Point returned incorrect M: %f\n", point.M())
	}

	if point.Stride() != 3 {
		t.Errorf("Point with dim XY returned stride %d instead of 3", point.Stride())
	}
}

func TestXYMPoint(t *testing.T) {
	t.Parallel()

	point := geom.NewPoint(geom.XYM, []float64{1.0, 2.0, 3.0})

	if point.Empty() {
		t.Error("Point with data {1.0, 2.0, 3.0} returned Empty()\n")
	}

	if point.X() != 1.0 {
		t.Errorf("Point returned incorrect X: %f\n", point.X())
	}

	if point.Y() != 2.0 {
		t.Errorf("Point returned incorrect Y: %f\n", point.Y())
	}

	if !math.IsNaN(point.Z()) {
		t.Errorf("Point returned incorrect Z: %f\n", point.Z())
	}

	if point.M() != 3.0 {
		t.Errorf("Point returned incorrect M: %f\n", point.M())
	}

	if point.Stride() != 3 {
		t.Errorf("Point with dim XY returned stride %d instead of 3", point.Stride())
	}
}

func TestXYZMPoint(t *testing.T) {
	t.Parallel()

	point := geom.NewPoint(geom.XYZM, []float64{1.0, 2.0, 3.0, 4.0})

	if point.Empty() {
		t.Error("Point with data {1.0, 2.0, 3.0, 4.0} returned Empty()\n")
	}

	if point.X() != 1.0 {
		t.Errorf("Point returned incorrect X: %f\n", point.X())
	}

	if point.Y() != 2.0 {
		t.Errorf("Point returned incorrect Y: %f\n", point.Y())
	}

	if point.Z() != 3.0 {
		t.Errorf("Point returned incorrect Z: %f\n", point.Z())
	}

	if point.M() != 4.0 {
		t.Errorf("Point returned incorrect M: %f\n", point.M())
	}

	if point.Stride() != 4 {
		t.Errorf("Point with dim XY returned stride %d instead of 4", point.Stride())
	}
}
