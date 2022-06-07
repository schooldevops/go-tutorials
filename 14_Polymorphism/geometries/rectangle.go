package geometries

type Rectangle struct {
	Name string
	X    float64
	Y    float64
}

func (rectangle Rectangle) GetName() string {
	return rectangle.Name
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.X * rectangle.Y
}
