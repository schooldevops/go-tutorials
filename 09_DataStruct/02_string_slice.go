package main

import "fmt"

func main() {
	greeting := "Hello World"
	fmt.Printf("Slice [1:5] of %s is [%s]\n", greeting, greeting[1:5])
	fmt.Printf("Index 1 : str(%s), Hex(%x), Ch(%c)\n", greeting[1:2], greeting[1], greeting[1])
	fmt.Println("just print: ", greeting[1])
}
