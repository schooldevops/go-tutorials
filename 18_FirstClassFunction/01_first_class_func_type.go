package main

import "fmt"

type add func(val01 int, val02 int) int

func main() {
	var funcVal add = func(val01 int, val02 int) int {
		return val01 + val02
	}

	result := funcVal(10, 20)
	fmt.Println("Result is:", result)
}
