package main

import "fmt"

type add func(val01 int, val02 int) int

func getAddFunc() add {
	var funcVal add = func(val01 int, val02 int) int {
		return val01 + val02
	}

	return funcVal
}

func main() {

	add2 := getAddFunc()
	result := add2(10, 20)
	fmt.Println("Result is:", result)
}
