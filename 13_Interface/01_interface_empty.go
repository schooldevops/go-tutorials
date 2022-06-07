package main

import "fmt"

type Students struct {
	name    string
	korean  int
	english int
	math    int
}

func showInterface(i interface{}) {
	fmt.Printf("Interface type %T, value is %v\n", i, i)
}

func main() {
	str := "Hello World"
	showInterface(str)

	age := 28
	showInterface(age)

	kido := Students{
		name:    "Kido",
		korean:  95,
		english: 90,
		math:    80,
	}

	showInterface(kido)

}
