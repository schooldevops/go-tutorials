package main

import "fmt"

func main() {
	a := [...]string{"Go", "Java", "Python"}
	b := a
	b[0] = "Kotlin"
	fmt.Println("string a: ", a)
	fmt.Println("string b: ", b)
}
