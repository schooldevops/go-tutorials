package main

import (
	"oop_new/students"
)

func main() {
	s := students.New("Kido", 95, 90, 80)
	s.TotalScore()
}
