package main

import "fmt"

func main() {

	userScore := make(map[string]int)
	userScore["Kido"] = 90
	userScore["Musk"] = 99
	userScore["Mario"] = 60
	fmt.Println("User Score Info: ", userScore)
}
