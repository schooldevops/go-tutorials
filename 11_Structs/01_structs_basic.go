package main

import (
	"fmt"
)

type Student struct {
	name           string
	korean         int
	english        int
	math           int
	weight, height float32
}

func main() {
	kido := Student{
		name:    "Kido",
		korean:  80,
		english: 70,
		math:    90,
		weight:  80,
		height:  177,
	}

	mario := Student{"Mario", 60, 90, 88, 75, 160}

	fmt.Println("Student kido is", kido)
	fmt.Println("Student mario is", mario)
}
