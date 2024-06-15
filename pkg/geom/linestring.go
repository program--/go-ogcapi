package geom

import "fmt"

type LineString struct {
	coordsD1
}

func NewLineString(dims DimsType, data []float64) *LineString {
	line := new(LineString)
	line.data = data
	line.dims = dims
	return line
}

func (l *LineString) Type() GeomType {
	return LINESTRING
}

func (l *LineString) PointAt(idx int) (*Point, error) {
	if l.data == nil || idx >= l.NumPoints() {
		return nil, fmt.Errorf("LineString.PointAt(%d) out of range of %d", idx, l.NumPoints())
	}

	start := idx * l.Stride()
	end := start + l.Stride()
	return NewPoint(l.dims, l.data[start:end]), nil
}

func (l *LineString) NumPoints() int {
	return len(l.data) / l.Stride()
}
