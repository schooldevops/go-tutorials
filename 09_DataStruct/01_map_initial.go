package main

import "fmt"

func main() {
	userScore := map[string]int{
		"Kido":  90,
		"Musk":  99,
		"Mario": 60,
	}
	fmt.Println("User Score Info: ", userScore)
	fmt.Println("Total length of map", len(userScore))

}
