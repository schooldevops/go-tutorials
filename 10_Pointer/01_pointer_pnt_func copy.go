package main

import "fmt"

func pnt_function() *int {
	number := 255
	return &number
}

func pnt_function_v2(num *int) *int {
	return num
}

func main() {
	pnt_func := pnt_function()
	fmt.Println("Value of pnt_func is", *pnt_func)
	*pnt_func = 100
	fmt.Println("Value of pnt_func is", *pnt_func)

	number := 100
	pnt_func_v2 := pnt_function_v2(&number)
	fmt.Println("number is", number)
	fmt.Println("pnt_func_v2 is", *pnt_func_v2)
	*pnt_func_v2 = 150
	fmt.Println("pnt_func_v2 is", *pnt_func_v2)
	fmt.Println("number is", number)
}
