package main

import "oop/students"

func main() {
	kido := students.Student{
		Name:    "Kido",
		Korean:  95,
		English: 90,
		Math:    80,
	}

	kido.TotalScore()
}
