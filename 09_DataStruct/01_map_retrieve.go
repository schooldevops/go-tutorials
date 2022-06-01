package main

import "fmt"

func main() {
	userScore := map[string]int{
		"Kido":  90,
		"Musk":  99,
		"Mario": 60,
	}

	user := "Kido"
	score := userScore[user]
	fmt.Printf("User %s, Score is %d\n", user, score)

	user = "Steve"
	score = userScore[user]
	fmt.Printf("User %s, Score is %d\n", user, score)
}
