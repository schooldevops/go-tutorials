package main

import (
	"fmt"
)

func main() {
	dice := 3

	if dice%2 == 0 {
		fmt.Println("Dice number", dice, "is event.")
	} else {
		fmt.Println("Dice number", dice, "is odd.")
	}
}
