package main

import (
	"fmt"
)

func main() {
	language := [][]string{
		{"Java", "Spring"},
		{"Kotlin", "Spring"},
		{"Python", "Django"},
	}

	for _, row := range language {
		for _, col := range row {
			fmt.Printf("%s ", col)
		}
		fmt.Println()
	}
}
