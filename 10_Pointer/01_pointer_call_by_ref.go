package main

import "fmt"

func changeValue(pnt_value *int, val int) {
	*pnt_value = val
}

func main() {
	number := 128
	fmt.Println("number is ", number)
	pnt_number := &number
	changeValue(pnt_number, 50)
	fmt.Println("Changed value is ", number)
}
