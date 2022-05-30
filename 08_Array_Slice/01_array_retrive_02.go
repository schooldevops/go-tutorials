package main

import "fmt"

func main() {
	scores := [...]int{40, 50, 80, 79, 95, 100}
	fmt.Println("Length of scores :", len(scores))

	for i, value := range scores {
		fmt.Printf("Index [%d], Value [%d]\n", i, value)
	}
}
