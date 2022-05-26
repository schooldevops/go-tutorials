package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("i is 5. breaking loop")
			break
		}
		fmt.Println("Current Value", i)
	}

}
