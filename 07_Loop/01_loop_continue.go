package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println()
			continue
		}
		fmt.Println("Current Value", i)
	}
}
