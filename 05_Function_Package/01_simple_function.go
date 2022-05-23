package main

import (
	"fmt"
)

func triangleArea(width, height int) int {
	area := width * height / 2
	return area
}

func main() {
	width, height := 10, 5
	areaResult := triangleArea(width, height)

	fmt.Println("Triangle Area :", areaResult)
}
