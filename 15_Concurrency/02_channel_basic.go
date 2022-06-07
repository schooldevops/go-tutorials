package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int)

	fmt.Printf("Type of channel is %\n", ch)
}
