package main

import "fmt"

func main() {
	scores := [...]int{40, 50, 80, 79, 95, 100}
	fmt.Println("Length of scores :", len(scores))

	for i := 0; i < len(scores); i++ {
		fmt.Printf("Index [%d], Value [%d]\n", i, scores[i])
	}
}
