package geometries 

import "math"

type Circle struct {
	Name string
	R float64
}

func (circle Circle) GetName() string {
	return circle.Name
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.R * circle.R
}