package geom_test

import (
	"testing"

	"github.com/program--/go-ogcapi/pkg/geom"
)

func TestEmptyPolygon(t *testing.T) {
	t.Parallel()

	empty_poly := geom.NewPolygon(geom.XY, nil)

	if empty_poly.Type() != geom.POLYGON {
		t.Errorf("Polygon returned incorrect GeomType: %x\n", empty_poly.Type())
	}

	if !empty_poly.Empty() {
		t.Error("Empty polygon returned !Empty()\n")
	}

	if empty_poly.NumRings() > 0 {
		t.Errorf("Empty polygon returned %d NumRings()\n", empty_poly.NumRings())
	}

	if empty_poly.NumHoles() > 0 {
		t.Errorf("Empty polygon returned %d NumHoles()\n", empty_poly.NumHoles())
	}

	if ring, err := empty_poly.RingAt(0); err == nil {
		t.Errorf("Empty polygon returned non-nil ring: %v\n", ring.Coords())
	}
}
