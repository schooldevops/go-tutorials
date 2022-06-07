package main

import "fmt"

func main() {
	young := struct {
		name           string
		korean         int
		english        int
		math           int
		weight, height float32
	}{
		name:    "Mario",
		korean:  90,
		english: 89,
		math:    85,
		weight:  80,
		height:  177,
	}

	fmt.Println("Anonymous struct ", young)
}
