package main

import "fmt"

func main() {
	original_value := "Hello World!"
	fmt.Println("Original Value is ", original_value)
	for _, value := range []rune(original_value) {
		defer fmt.Printf("%c", value)
	}
}
