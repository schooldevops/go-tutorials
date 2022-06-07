package main

import (
	"fmt"
)

type Student struct {
	name    string
	korean  int
	english int
	math    int
}

func main() {
	kido := Student{
		name:    "Kido",
		korean:  95,
		english: 90,
		math:    85,
	}
	kido_clone := Student{
		name:    "Kido",
		korean:  95,
		english: 90,
		math:    85,
	}

	if kido == kido_clone {
		fmt.Println("kido and kido_clone is equals.")
	} else {
		fmt.Println("kido and kido_clone is not equals.")
	}

	kido_clone_v2 := Student{
		korean:  95,
		english: 90,
		math:    85,
	}

	if kido == kido_clone_v2 {
		fmt.Println("kido and kido_clone is equals.")
	} else {
		fmt.Println("kido and kido_clone_v2 is not equals.")
	}

	kido_clone_v2.name = "Kido"

	if kido == kido_clone_v2 {
		fmt.Println("kido and kido_clone is equals.")
	} else {
		fmt.Println("kido and kido_clone_v2 is not equals.")
	}
}
