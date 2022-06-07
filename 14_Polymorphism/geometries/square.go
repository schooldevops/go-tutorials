package geometries

type Square struct {
	Name string
	X    float64
}

func (square Square) GetName() string {
	return square.Name
}

func (square Square) Area() float64 {
	return square.X * square.X
}