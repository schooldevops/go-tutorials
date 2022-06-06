package main

import "fmt"

func main() {
	pnt_number := new(int)
	fmt.Printf("pnt_number value is %d, type is %T, address is %v\n", *pnt_number, pnt_number, pnt_number)
	*pnt_number = 255
	fmt.Println("pnt_number value is", *pnt_number)
}
