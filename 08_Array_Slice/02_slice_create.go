package main

import (
	"fmt"
)

func main() {
	scores := [5]int{80, 70, 95, 100, 97}
	var score_slice []int = scores[1:3]
	fmt.Println("score :", scores)
	fmt.Println("slice :", score_slice)
	score_slice[0] = 100
	fmt.Println("after score :", scores)
	fmt.Println("after slice :", score_slice)

}
