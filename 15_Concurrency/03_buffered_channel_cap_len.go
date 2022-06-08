package main

import (
	"fmt"
)

func main() {
	buf_ch := make(chan int, 5)
	buf_ch <- 1
	buf_ch <- 2
	buf_ch <- 3

	fmt.Println("capacity is", cap(buf_ch))
	fmt.Println("length is", len(buf_ch))

	fmt.Println("read from cahnnel", <-buf_ch)
	fmt.Println("capacity is", cap(buf_ch))
	fmt.Println("length is", len(buf_ch))
}
