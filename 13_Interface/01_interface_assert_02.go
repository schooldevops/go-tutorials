package main

import (
	"fmt"
)

func assertInteger(i interface{}) {
	value, ok := i.(int)
	fmt.Printf("Value is %v, is assert? %t\n", value, ok)
}

func main() {
	var value interface{} = 128
	assertInteger(value)

	var value2 interface{} = "Hello"
	assertInteger(value2)
}
