package geom

/*
Internal coordinates format is based heavily on
https://github.com/twpayne/go-geom, which is distributed
under a BSD-2-Clause license.
*/

// Zero-dimensional coordinates (Points)
type coordsD0 struct {
	dims DimsType
	data []float64
}

// One-dimensional coordinates (LineString, MultiPoint)
type coordsD1 struct {
	coordsD0
}

// Two-dimensional coordinates (Polygon, MultiLineString)
type coordsD2 struct {
	coordsD1

	// Offset array
	//
	// Each element of this holds
	// the ending index of a set of coordinates.
	offsets []int
}

// Three-dimensional coordinates (MultiPolygon)
type coordsD3 struct {
	coordsD1

	// Jagged offset array
	//
	// Each element (i.e. each []int) of this
	// holds the ring offsets for a polygon.
	offsets [][]int
}

// Return a slice of the flat coordinates.
func (c *coordsD0) Coords() []float64 {
	return c.data
}

// Returns true if the backing data array is nil or has no elements.
func (c *coordsD0) Empty() bool {
	return c.data == nil || len(c.data) == 0
}

// Returns the coordinate dimensions (not the shape dimension).
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

//lint:ignore U1000 Ignore unused function temporarily
func inflateD2(data []float64, offsets []int) [][]float64 {
	coords := make([][]float64, len(offsets))

	start := 0
	for i, offset := range offsets {
		coords[i] = data[start:offset]
		start = offset
	}
	return coords
}

//lint:ignore U1000 Ignore unused function temporarily
func inflateD3(data []float64, offsets [][]int) [][][]float64 {
	coords := make([][][]float64, len(offsets))
	start := 0
	for i, goffset := range offsets {
		end := goffset[len(goffset)-1]
		coords[i] = inflateD2(data[start:end], goffset)
		start = end
	}

	return coords
}

func deflateD2(coords [][]float64) ([]float64, []int) {
	data := make([]float64, 0)
	offsets := make([]int, 0, len(coords))
	for _, slice := range coords {
		offsets = append(offsets, len(slice))
		data = append(data, slice...)
	}

	return data, offsets
}

func deflateD3(coords [][][]float64) ([]float64, [][]int) {
	data := make([]float64, 0)
	offsets := make([][]int, len(coords))

	var slice []float64 = nil
	for i, nested := range coords {
		slice, offsets[i] = deflateD2(nested)
		data = append(data, slice...)
	}

	return data, offsets
}
