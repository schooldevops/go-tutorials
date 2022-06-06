package main

import "fmt"

func passArrayLoc(arr *[3]int) {
	(*arr)[0] = 100
}

func passSlice(slice []int) {
	slice[0] = 100
}

func main() {
	scores := [3]int{75, 80, 99}
	passArrayLoc(&scores)
	fmt.Println("Scores :", scores)

	scores_v2 := []int{75, 80, 90}
	passSlice(scores_v2)
	fmt.Println("Score v2 :", scores_v2)
}
