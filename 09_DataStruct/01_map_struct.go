package main

import "fmt"

type student struct {
	name    string
	korean  int
	english int
	math    int
}

func main() {
	kido := student{
		name:    "Kido",
		korean:  80,
		english: 85,
		math:    90,
	}

	musk := student{
		name:    "Musk",
		korean:  60,
		english: 99,
		math:    100,
	}

	mario := student{
		name:    "Mario",
		korean:  70,
		english: 90,
		math:    95,
	}

	students := map[string]student{
		"Kido":  kido,
		"Musk":  musk,
		"Mario": mario,
	}

	for key, value := range students {
		fmt.Printf("User key is %s, Korean is %d, English is %d, Math is %d\n", key, value.korean, value.english, value.math)
	}
}
