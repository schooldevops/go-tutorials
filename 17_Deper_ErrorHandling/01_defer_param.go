package main

import "fmt"

func defer_statement(param int) {
	fmt.Println("Defer Statement start with param:", param)
}

func main() {
	p := 100
	defer defer_statement(p)
	p = 200
	fmt.Println("P value is ", p)
}
