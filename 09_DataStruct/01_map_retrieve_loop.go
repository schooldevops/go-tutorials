package main

import "fmt"

func main() {
	userScore := map[string]int{
		"Kido":  90,
		"Musk":  99,
		"Mario": 60,
	}

	for key, value := range userScore {
		fmt.Printf("Key is %s, Value is %d\n", key, value)
	}
}
