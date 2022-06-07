package main

import "fmt"

type Student struct {
	name    string
	korean  int
	english int
	math    int
}

func main() {
	var kido Student
	fmt.Println("Name is", kido.name)
	fmt.Println("Korean Score is", kido.korean)
	fmt.Println("English Score is", kido.english)
	fmt.Println("Math Score is", kido.math)

	kido.korean = 99
	fmt.Println("Korean Score is", kido.korean)
}
