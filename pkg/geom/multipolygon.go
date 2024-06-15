package geom

import "fmt"

// TODO: Tests
type MultiPolygon struct {
	coordsD3
}

func NewMultiPolygon(dims DimsType, data [][][]float64) *MultiPolygon {
	mp := new(MultiPolygon)
	mp.data, mp.offsets = deflateD3(data)
	return mp
}

func (*MultiPolygon) Type() GeomType {
	return MULTIPOLYGON
}

func (m *MultiPolygon) NumPolys() int {
	return len(m.offsets)
}

func (m *MultiPolygon) NumPoints() int {
	return len(m.data) / m.Stride()
}

func (m *MultiPolygon) PolyAt(idx int) (*Polygon, error) {
	if m.Empty() || idx >= m.NumPolys() {
		return nil, fmt.Errorf("MultiPolygon.PolyAt(%d) out of range of %d", idx, m.NumPolys())
	}

	var start int
	if idx == 0 {
		start = 0
	} else {
		start = m.offsets[idx-1][len(m.offsets[idx-1])-1]
	}

	end := m.offsets[idx][len(m.offsets[idx])-1]
	return NewPolygonFlat(m.dims, m.data[start:end], m.offsets[idx]), nil
}
