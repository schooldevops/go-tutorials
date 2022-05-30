package main

import (
	"fmt"
)

func main() {
	language := []string{"Java", "Go", "Python"}
	fmt.Println("1. Language Slice", language, "length", len(language), "capacity", cap(language))
	language = append(language, "Kotlin")
	fmt.Println("2. Language Slice", language, "length", len(language), "capacity", cap(language))
	language = append(language, "Algol", "Fortran")
	fmt.Println("3. Language Slice", language, "length", len(language), "capacity", cap(language))
	myFavorite := []string{"C", "C++", "Swift", "JavaScript"}
	language = append(language, myFavorite...)
	fmt.Println("4. Language Slice", language, "length", len(language), "capacity", cap(language))
}
