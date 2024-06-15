package geom

import "fmt"

type MultiLineString struct {
	coordsD2
}

func NewMultiLineString(dims DimsType, data [][]float64) *MultiLineString {
	ml := new(MultiLineString)
	ml.data = make([]float64, 0)
	ml.offsets = make([]int, 0, len(data))
	for _, slice := range data {
		ml.offsets = append(ml.offsets, len(slice))
		ml.data = append(ml.data, slice...)
	}

	return ml
}

func (*MultiLineString) Type() GeomType {
	return MULTILINESTRING
}

func (m *MultiLineString) LineAt(idx int) (*LineString, error) {
	if m.Empty() {
		return nil, fmt.Errorf("MultiLineString.LineAt(%d) out of range of %d", idx, m.NumLines())
	}

	var start int
	if idx == 0 {
		start = 0
	} else {
		start = m.offsets[idx-1]
	}
	end := m.offsets[idx]

	return NewLineString(m.dims, m.data[start:end]), nil
}

func (m *MultiLineString) NumLines() int {
	return len(m.offsets)
}

func (m *MultiLineString) NumPoints() int {
	return len(m.data) / m.Stride()
}
