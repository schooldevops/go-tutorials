package main

import (
	"fmt"
)

func main() {
	cube := make([]int, 5, 10)
	fmt.Println(cube)
	fmt.Printf("Slice length is %d, Slice capacity is %d\n", len(cube), cap(cube))

}
