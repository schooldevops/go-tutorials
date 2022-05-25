package main

import (
	"fmt"
)

func main() {
	dice := 4

	switch {
	case dice == 1:
		fmt.Printf("Move %d step\n", dice)
	case dice >= 2 && dice <= 6:
		fmt.Printf("Move %d steps\n", dice)
	default:
		fmt.Println("Value Error...")
	}
}
