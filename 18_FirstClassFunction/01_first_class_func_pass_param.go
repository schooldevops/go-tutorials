package main

import "fmt"

type add2 func(param1 int, param2 int) int

func printAddResult(add func(param1 int, param2 int) int) {
	fmt.Println("funcParam: ", (add(10, 20)))
}

func printAddResultWithType(add add2) {
	fmt.Println("funcParam with Type: ", (add(10, 20)))
}

func main() {
	add := func(param1 int, param2 int) int {
		return param1 + param2
	}

	printAddResult(add)
	printAddResultWithType(add)
}
