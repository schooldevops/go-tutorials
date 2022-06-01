package main

import "fmt"

func main() {
	userScore := map[string]int{
		"Kido":  90,
		"Musk":  99,
		"Mario": 60,
	}

	fmt.Println("Map data is", userScore)
	delete(userScore, "Kido")
	fmt.Println("Map data after deletion", userScore)
}
