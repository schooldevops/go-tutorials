package main

import "fmt"

type Student struct {
	name    string
	korean  int
	english int
	math    int
}

func main() {
	kido := Student{"Kido", 95, 90, 85}

	fmt.Println("Name is", kido.name)
	fmt.Println("Korean Score is", kido.korean)
	fmt.Println("English Score is", kido.english)
	fmt.Println("Math Score is", kido.math)
}
