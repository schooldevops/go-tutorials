package main

import (
	"fmt"
)

type Students struct {
	name    string
	korean  int
	english int
	math    int
}

func (s Students) totalAverageScore() {
	totalAvg := (float32)(s.korean+s.english+s.math) / 3.0
	fmt.Printf("Student %s, korean score is %d, english score is %d, math score is%d\n", s.name, s.korean, s.english, s.math)
	fmt.Printf("Total Average is %f\n", totalAvg)
}

func main() {
	kido := Students{
		name:    "Kido",
		korean:  95,
		english: 90,
		math:    80,
	}

	kido.totalAverageScore()
}
