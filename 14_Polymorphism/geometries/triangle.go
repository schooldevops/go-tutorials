package geometries

type Triangle struct {
	Name string
	X    float64
	Y    float64
}

func (triangle Triangle) GetName() string {
	return triangle.Name
}

func (triangle Triangle) Area() float64 {
	return triangle.X * triangle.Y / 2
}