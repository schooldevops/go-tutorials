package main

import (
	"fmt"
	"reflect"
)

func main() {
	var count = 100
	fmt.Println("Count type is ", reflect.TypeOf(count))
}
