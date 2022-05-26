package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		fmt.Printf("---- %d단 ----\n", i)

		for j := 1; j <= 9; j++ {
			fmt.Printf(" %d x %d = %d\n", i, j, i*j)
		}

		fmt.Println()
	}
}
