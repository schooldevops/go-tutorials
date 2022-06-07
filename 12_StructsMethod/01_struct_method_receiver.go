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

func (s Students) changeName(newName string) {
	s.name = newName
	fmt.Println("Inside changeName:", s)
}

func (s *Students) changeKorean(score int) {
	s.korean = score
	fmt.Println("Inside changeKorean:", s)
}

func main() {
	kido := Students{
		name:    "Kido",
		korean:  60,
		english: 90,
		math:    80,
	}

	fmt.Println("1. origin value", kido)
	kido.changeName("Mario")
	fmt.Println("2. after changeName", kido)
	(&kido).changeKorean(100)
	fmt.Println("3. after changeKorean", kido)
}
