package main

import "fmt"

func changeValue(language [3]string) {
	language[0] = "Kotlin"
	fmt.Println("Change Value :", language)
}

func main() {
	lan := [...]string{"Go", "Java", "Python"}
	fmt.Println("Original Value :", lan)
	changeValue(lan)
	fmt.Println("After Value :", lan)
}
