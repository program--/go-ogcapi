package geom

type coordsD0 struct {
	dims DimsType
	data []float64
}

type coordsD1 struct {
	coordsD0
}

type coordsD2 struct {
	coordsD1
	offsets []int
}

type coordsD3 struct {
	coordsD1
	offsets [][]int
}

func (c *coordsD0) Coords() []float64 {
	return c.data
}

func (c *coordsD0) Empty() bool {
	return c.data == nil || len(c.data) == 0
}

func (c *coordsD0) Dimensions() DimsType {
	return c.dims
}

func (c *coordsD0) indexZ() int {
	switch c.dims {
	case XY:
		fallthrough
	case XYM:
		fallthrough
	default:
		return -1

	case XYZ:
		fallthrough
	case XYZM:
		return 2
	}
}

func (c *coordsD0) indexM() int {
	switch c.dims {
	case XY:
		fallthrough
	case XYZ:
		fallthrough
	default:
		return -1

	case XYM:
		return 2

	case XYZM:
		return 3
	}
}

func (c *coordsD0) Stride() int {
	switch c.dims {
	case XY:
		fallthrough
	default:
		return 2

	case XYZ:
		fallthrough
	case XYM:
		return 3

	case XYZM:
		return 4
	}
}
