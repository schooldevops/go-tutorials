package main

import (
	"fmt"
	"math/rand"
	"time"
)

func throwADice() int {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	return dice
}

func main() {

	switch dice := throwADice(); {
	case dice < 3:
		fmt.Printf("%d is lesser than 3\n", dice)
		fallthrough
	case dice < 5:
		fmt.Printf("%d is lesser than 5\n", dice)
		fallthrough
	case dice <= 6:
		fmt.Printf("%d is lesser than equals 6\n", dice)
	}
}
