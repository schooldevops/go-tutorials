package main

import (
	"fmt"
)

func rectangleInfo(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = (width + height) * 2

	return
}

func main() {
	width, height := 10, 5
	area, perimeter := rectangleInfo(width, height)

	fmt.Printf("Area is %d, Perimeter is %d", area, perimeter)
}
