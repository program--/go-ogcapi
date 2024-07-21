package geom

import (
	"fmt"
	"math"
)

type BoundingBox struct {
	coordsD1
}

func NewBoundingBox(dims DimsType, data []float64) (*BoundingBox, error) {
	bbox := new(BoundingBox)

	if len(data) != 2*bbox.Stride() {
		return nil, fmt.Errorf(
			"could not create new bounding box: data length mismatch %d != %d",
			len(data),
			2*bbox.Stride(),
		)
	}

	bbox.data = data
	bbox.dims = dims
	return bbox, nil
}

func (bbox *BoundingBox) XMin() float64 {
	return bbox.data[0]
}

func (bbox *BoundingBox) XMax() float64 {
	return bbox.data[0+bbox.Stride()]
}

func (bbox *BoundingBox) YMin() float64 {
	return bbox.data[1]
}

func (bbox *BoundingBox) YMax() float64 {
	return bbox.data[1+bbox.Stride()]
}

func (bbox *BoundingBox) ZMin() float64 {
	switch bbox.dims {
	case XYZ:
		fallthrough
	case XYZM:
		return bbox.data[2]

	case XY:
		fallthrough
	case XYM:
		fallthrough
	default:
		return math.NaN()
	}
}

func (bbox *BoundingBox) ZMax() float64 {
	switch bbox.dims {
	case XYZ:
		fallthrough
	case XYZM:
		return bbox.data[2+bbox.Stride()]

	case XY:
		fallthrough
	case XYM:
		fallthrough
	default:
		return math.NaN()
	}
}

func (bbox *BoundingBox) MMin() float64 {
	switch bbox.dims {
	case XYM:
		fallthrough
	case XYZM:
		return bbox.data[3]

	case XY:
		fallthrough
	case XYZ:
		fallthrough
	default:
		return math.NaN()
	}
}

func (bbox *BoundingBox) MMax() float64 {
	switch bbox.dims {
	case XYM:
		fallthrough
	case XYZM:
		return bbox.data[3+bbox.Stride()]

	case XY:
		fallthrough
	case XYZ:
		fallthrough
	default:
		return math.NaN()
	}
}
