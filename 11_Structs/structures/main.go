package main

import (
	"fmt"
	"structures/students"
)

func main() {
	kido := students.Students{
		Name:    "Kido",
		Korean:  95,
		English: 90,
		Math:    80,
	}

	fmt.Println("Student is", kido)
}
