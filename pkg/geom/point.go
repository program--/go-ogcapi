package geom

import "math"

type Point struct {
	coordsD0
}

func NewPoint(dims DimsType, data []float64) *Point {
	pt := new(Point)
	pt.data = data
	pt.dims = dims
	return pt
}

func (p *Point) Type() GeomType {
	return POINT
}

func (p *Point) X() float64 {
	return p.data[0]
}

func (p *Point) Y() float64 {
	return p.data[1]
}

func (p *Point) Z() float64 {
	idx := p.indexZ()
	if idx < 0 {
		return math.NaN()
	}

	return p.data[idx]
}

func (p *Point) M() float64 {
	idx := p.indexM()
	if idx < 0 {
		return math.NaN()
	}

	return p.data[idx]
}
