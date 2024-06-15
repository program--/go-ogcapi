package geom

import "fmt"

type MultiPoint struct {
	coordsD1
}

func NewMultiPoint(dims DimsType, data []float64) *MultiPoint {
	mp := new(MultiPoint)
	mp.data = data
	mp.dims = dims
	return mp
}

func (*MultiPoint) Type() GeomType {
	return MULTIPOINT
}

func (m *MultiPoint) NumPoints() int {
	return len(m.data) / m.Stride()
}

func (m *MultiPoint) PointAt(idx int) (*Point, error) {
	if m.data == nil || idx >= m.NumPoints() {
		return nil, fmt.Errorf("MultiPoint.PointAt(%d) out of range of %d", idx, m.NumPoints())
	}

	start := idx * m.Stride()
	end := start + m.Stride()
	return NewPoint(m.dims, m.data[start:end]), nil
}
