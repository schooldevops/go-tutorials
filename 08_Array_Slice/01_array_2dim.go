package main

import (
	"fmt"
)

func main() {

	languages := [3][2]string{
		{"Java", "Spring"},
		{"Go", "Gin"},
		{"Python", "FastAPI"},
	}

	fmt.Println(languages)
	fmt.Println("Len :", len(languages))
}
