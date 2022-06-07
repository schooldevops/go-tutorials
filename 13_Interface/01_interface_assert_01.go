package main

import (
	"fmt"
)

func assertInteger(i interface{}) {
	value := i.(int)
	fmt.Println("Value is", value)
}

func main() {
	var value interface{} = 128
	assertInteger(value)

	var value2 interface{} = "Hello"
	assertInteger(value2)
}
