package geom

import "fmt"

type MultiPolygon struct {
	coordsD3
}

func NewMultiPolygon(dims DimsType, data [][][]float64) *MultiPolygon {
	mp := new(MultiPolygon)
	mp.data = make([]float64, 0)
	mp.offsets = make([][]int, len(data))

	for i, nested := range data {
		mp.offsets[i] = make([]int, 0, len(nested))

		for _, slice := range nested {
			mp.offsets[i] = append(mp.offsets[i], len(slice))
			mp.data = append(mp.data, slice...)
		}
	}

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
	if m.Empty() {
		return nil, fmt.Errorf("MultiPolygon.PolyAt(%d) out of range of %d", idx, m.NumPolys())
	}

	// TODO: need new geometry from flat coords + offsets func
	return nil, nil
}
