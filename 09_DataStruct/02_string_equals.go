package main

import "fmt"

func main() {
	string1 := "Go Lang"
	string2 := "Go Lang"
	fmt.Printf("Are %s and %s equal? --> %t\n", string1, string2, string1 == string2)

	string3 := "Lang Go"
	fmt.Printf("Are %s and %s equal? --> %t\n", string2, string3, string2 == string3)
}
