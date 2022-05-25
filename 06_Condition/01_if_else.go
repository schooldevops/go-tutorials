package main

import (
	"fmt"
)

func main() {
	dice := 4

	if dice > 6 && dice < 1 {
		fmt.Println("Dice Number never be over 6 and under 1")
	} else if dice%2 == 0 {
		fmt.Println("Dice number", dice, "is event.")
	} else {
		fmt.Println("Dice number", dice, "is odd.")
	}
}
