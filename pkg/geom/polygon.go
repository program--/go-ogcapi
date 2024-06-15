package geom

import "fmt"

// TODO: Tests
type Polygon struct {
	coordsD2
}

func NewPolygon(dims DimsType, data [][]float64) *Polygon {
	poly := new(Polygon)
	poly.data, poly.offsets = deflateD2(data)
	return poly
}

func NewPolygonFlat(dims DimsType, data []float64, offsets []int) *Polygon {
	poly := new(Polygon)
	poly.data, poly.offsets = data, offsets
	return poly
}

func (*Polygon) Type() GeomType {
	return POLYGON
}

func (p *Polygon) RingAt(idx int) (*LineString, error) {
	if p.Empty() {
		return nil, fmt.Errorf("Polygon.RingAt(%d) out of range of %d", idx, p.NumRings())
	}

	var start, end int

	if idx == 0 {
		start = 0
	} else {
		start = p.offsets[idx-1]
	}

	end = p.offsets[idx]

	return NewLineString(p.dims, p.data[start:end]), nil
}

func (p *Polygon) NumRings() int {
	return len(p.offsets)
}

func (p *Polygon) NumHoles() int {
	if p.Empty() {
		return 0
	}

	return len(p.offsets) - 1
}

func (p *Polygon) NumPoints() int {
	return len(p.data) / p.Stride()
}
