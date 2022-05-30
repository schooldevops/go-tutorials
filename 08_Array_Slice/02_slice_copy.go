package main

import (
	"fmt"
)

func main() {
	languages := []string{"Java", "Go", "Python", "Swift", "JavaScript", "Ruby", "Gradle"}
	slicedLanguage := languages[0:5]
	fmt.Println("1. Language:", languages, len(languages), cap(languages))
	fmt.Println("2. Language:", slicedLanguage, len(slicedLanguage), cap(slicedLanguage))
	fmt.Printf("3. orig address: %p, sliced address: %p\n", &languages[0], &slicedLanguage[0])

	copyedLanguage := make([]string, len(slicedLanguage))
	copy(copyedLanguage, slicedLanguage)
	fmt.Println("4. Language:", copyedLanguage, len(copyedLanguage), cap(copyedLanguage))
	fmt.Printf("5. sliced address: %p, copyed address: %p\n", &slicedLanguage[0], &copyedLanguage[0])
}
