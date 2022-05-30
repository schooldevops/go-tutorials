package main

import (
	"fmt"
)

func changeLanguage(language []string, postfix string) {
	for i, value := range language {
		language[i] = value + postfix
	}
}

func main() {
	language := []string{"Java", "Go", "Python"}
	fmt.Println("1. Language is :", language)
	changeLanguage(language, " Language")
	fmt.Println("2. Language is :", language)
}
