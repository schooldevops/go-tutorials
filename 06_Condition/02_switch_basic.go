package main

import (
	"fmt"
)

func main() {
	dice := 4

	fmt.Println("Dice value is", dice)
	switch dice {
	case 1:
		fmt.Println("Move one step")
	case 2:
		fmt.Println("Move two steps")
	case 3:
		fmt.Println("Move three steps")
	case 4:
		fmt.Println("Move four steps")
	case 5:
		fmt.Println("Move five steps")
	case 6:
		fmt.Println("Move six steps")
	default:
		fmt.Println("Value Error...")
	}
}
