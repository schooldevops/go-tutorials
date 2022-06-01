package main

import "fmt"

func main() {
	userScore := map[string]int{
		"Kido":  90,
		"Musk":  99,
		"Mario": 60,
	}

	user := "Steve"
	score, ok := userScore[user]
	if ok == true {
		fmt.Printf("User %s, Score is %d\n", user, score)
	} else {
		fmt.Printf("User %s is not exists.\n", user)
	}
}
