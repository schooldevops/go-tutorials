package main

import (
	"fmt"
	"school/students"
)

func main() {
	bodyInfo := students.BodyInfo{
		Weight: 82,
		Height: 177,
	}

	kido := students.Student{
		Name:     "Kido",
		Korean:   95,
		English:  90,
		Math:     80,
		BodyInfo: bodyInfo,
	}

	fmt.Println("Student Name: ", kido.Name)
	fmt.Println("Total Score: ", kido.TotalScore())
	fmt.Println("Bmi:", kido.Bmi())
}
