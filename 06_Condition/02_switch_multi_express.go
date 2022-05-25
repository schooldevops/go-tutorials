package main

import (
	"fmt"
)

func main() {
	dice := 4

	fmt.Println("Dice value is", dice)
	switch dice {
	case 1:
		fmt.Printf("Move %d step\n", dice)
	case 2, 3, 4, 5, 6:
		fmt.Printf("Move %d steps\n", dice)
	default:
		fmt.Println("Value Error...")
	}
}
