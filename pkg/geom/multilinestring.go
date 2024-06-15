package geom

import "fmt"

// TODO: Tests
type MultiLineString struct {
	coordsD2
}

func NewMultiLineString(dims DimsType, data [][]float64) *MultiLineString {
	ml := new(MultiLineString)
	ml.data, ml.offsets = deflateD2(data)
	return ml
}

func NewMultiLineStringFlat(dims DimsType, data []float64, offsets []int) *MultiLineString {
	ml := new(MultiLineString)
	ml.data, ml.offsets = data, offsets
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
