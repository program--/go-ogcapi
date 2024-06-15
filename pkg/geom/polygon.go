package geom

type Polygon struct {
	coordsD2
}

func NewPolygon(dims DimsType, data [][]float64) *Polygon {
	poly := new(Polygon)
	poly.data = make([]float64, 0)
	poly.offsets = make([]int, len(data))
	for i, slice := range data {
		poly.offsets[i] = len(slice)
		poly.data = append(poly.data, slice...)
	}

	return poly
}
func (p *Polygon) Type() GeomType {
	return POLYGON
}

func (p *Polygon) RingAt(idx int) *LineString {
	var start, end int

	if idx == 0 {
		start = 0
	} else {
		start = p.offsets[idx-1]
	}

	end = p.offsets[idx]

	return NewLineString(p.dims, p.data[start:end])
}

func (p *Polygon) NumRings() int {
	return len(p.offsets)
}

func (p *Polygon) NumPoints() int {
	return len(p.data) / p.Stride()
}
