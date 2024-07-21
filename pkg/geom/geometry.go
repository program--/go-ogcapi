package geom

type GeomType uint8

const (
	GEOMETRY GeomType = iota
	POINT
	LINESTRING
	POLYGON
	MULTIPOINT
	MULTILINESTRING
	MULTIPOLYGON
	GEOMETRYCOLLECTION
)

type DimsType uint8

const (
	XY DimsType = iota
	XYZ
	XYM
	XYZM
)

func GeometryTypeToCode(geomType GeomType, dimsType DimsType) uint16 {
	return uint16(geomType) + (uint16(1000) * uint16(dimsType))
}

func GeometryTypeFromCode(code uint16) (GeomType, DimsType) {
	switch {
	case code < 1000:
		return GeomType(code), XY
	case code < 2000:
		return GeomType(code - 1000), XYZ
	case code < 3000:
		return GeomType(code - 2000), XYM
	default:
		return GeomType(code - 3000), XYZM
	}
}

type Geometry interface {
	Type() GeomType
	Dimensions() DimsType
	Coords() []float64
	Empty() bool
	Stride() int
}
