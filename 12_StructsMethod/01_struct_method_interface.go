package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

type Square struct {
	width int
}

type Triangle struct {
	width  int
	height int
}

type Circle struct {
	radius int
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func (s Square) Area() int {
	return s.width * s.width
}

func (t Triangle) Area() float32 {
	return (float32)(t.width*t.height) / 2.0
}

func (c Circle) Area() float32 {
	return 3.1415 * (float32)(c.radius) / 2.0
}

func main() {
	rect := Rectangle{
		width:  10,
		height: 5,
	}

	square := Square{
		width: 10,
	}

	triangle := Triangle{
		width:  10,
		height: 5,
	}

	circle := Circle{
		radius: 10,
	}

	fmt.Println("Rectangle Area is", rect.Area())
	fmt.Println("Square Area is", square.Area())
	fmt.Println("Triangle Area is", triangle.Area())
	fmt.Println("Circle Area is", circle.Area())
}
