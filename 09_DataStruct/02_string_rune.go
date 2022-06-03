package main

import "fmt"

func main() {
	greeting := "Hello World"

	fmt.Println(greeting)
	fmt.Println("Charancter with Rune:")
	runes := []rune(greeting)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()

	fmt.Println("Byte :")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%c ", greeting[i])
	}
	fmt.Println()
}
