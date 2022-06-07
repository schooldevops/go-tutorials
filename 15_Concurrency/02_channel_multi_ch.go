package main

import "fmt"

func sum(list []int, sum_ch chan int) {
	sum := 0
	for _, value := range list {
		sum += value
	}

	sum_ch <- sum
}

func multiplex(list []int, mul_ch chan int) {
	mul := 1

	for _, value := range list {
		mul *= value
	}

	mul_ch <- mul
}

func main() {
	sum_ch := make(chan int)
	mul_ch := make(chan int)

	value_list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go sum(value_list, sum_ch)
	go multiplex(value_list, mul_ch)

	total_sum, total_mul := <-sum_ch, <-mul_ch

	fmt.Printf("Total Sum is %d, Total Mul is %d\n", total_sum, total_mul)
}
