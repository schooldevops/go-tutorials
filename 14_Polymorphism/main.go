package main

import (
	geometries "polymorphism/geometries"
	"fmt"
)

func main() {

	circle := geometries.Circle{Name: "원", R: 10}
	rectangle := geometries.Rectangle{Name: "직사각형", X: 5, Y: 10}
	square := geometries.Square{Name: "정사각형", X: 5}
	triangle := geometries.Triangle{Name: "삼각형", X: 10, Y: 5}

	geometries := []geometries.Geometry{circle, rectangle, square, triangle}
	for _, geometry := range geometries {
		fmt.Println(geometry.GetName())
		fmt.Printf("Area: %0.2f\n", geometry.Area())
		fmt.Println("----------------------")
	}

}