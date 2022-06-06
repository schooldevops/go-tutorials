package main

import "fmt"

func main() {
	number := 128
	var pnt_number *int

	fmt.Println("Zero value of pnt_number is", pnt_number)
	pnt_number = &number
	fmt.Println("set pnt_number by address of number", pnt_number)
}
