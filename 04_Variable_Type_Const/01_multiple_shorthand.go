package main

import "fmt"

func main() {
	count := 10
	fmt.Println("Count =", count)

	name, age, score, pass := "kido", 40, 99, true
	fmt.Println("Student", name, "score is", score, "and his age is", age, "pass value is", pass)
}
