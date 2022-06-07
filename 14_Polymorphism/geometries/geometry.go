package geometries

type Geometry interface {
	GetName() string
	Area() float64
}