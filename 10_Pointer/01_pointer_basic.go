package main

import "fmt"

func main() {
	number := 128
	var pnt_number *int
	pnt_number = &number

	fmt.Printf("Number is %d, address is %p\n", number, &number)
	fmt.Println("pnt_nubmer is ", pnt_number)
}
