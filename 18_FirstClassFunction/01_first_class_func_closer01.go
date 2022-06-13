package main

import "fmt"

func main() {
	a := 5
	func() {
		fmt.Println("a = ", a)
		a = 10
		fmt.Println("a = ", a)
	}()
	fmt.Println("a = ", a)
}
