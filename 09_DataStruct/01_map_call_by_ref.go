package main

import "fmt"

func modifyScore(userScore map[string]int, key string, score int) {
	userScore[key] = score
}

func main() {

	userScore := map[string]int{
		"Kido":  90,
		"Musk":  99,
		"Mario": 60,
	}

	fmt.Println("UserScore ", userScore)
	modifyScore(userScore, "Kido", 100)
	fmt.Println("UserScore after modify ", userScore)
}
