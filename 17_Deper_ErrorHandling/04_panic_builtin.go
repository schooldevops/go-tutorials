package main

import "fmt"

func printSlice(paramSlice []int, number int) int {
	return paramSlice[number]
}

func main() {
	sl := []int{1, 2, 3, 4}
	fmt.Println("Result Value: ", printSlice(sl, 10))
}
