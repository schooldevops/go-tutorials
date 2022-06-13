package main

import "fmt"

func apply(s []int, applyFunc func(param int) int) []int {
	var resultSl []int

	for _, v := range s {
		resultSl = append(resultSl, applyFunc(v))
	}
	return resultSl
}

func main() {
	sliceVal := []int{1, 2, 3, 4, 5, 6, 7, 8}
	paramVal := 10

	addFunc := func(param int) int {
		return param + paramVal
	}

	mulFunc := func(param int) int {
		return param * paramVal
	}

	fmt.Println("apply results with addFunc:", apply(sliceVal, addFunc))
	fmt.Println("apply results with mulFunc:", apply(sliceVal, mulFunc))
}
