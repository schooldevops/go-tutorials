package main

import "fmt"

type Student struct {
	name    string
	korean  int
	english int
	math    int
}

func main() {
	pnt_kido := &Student{
		name:    "Kido",
		korean:  95,
		english: 90,
		math:    80,
	}

	fmt.Println("Name is", (*pnt_kido).name)
	fmt.Println("Korean is", (*pnt_kido).korean)
	fmt.Println("----------------------------------")
	fmt.Println("Name is", pnt_kido.name)
	fmt.Println("Korean is", pnt_kido.korean)
}
