package main

import "fmt"

func main() {
	buf_ch := make(chan string, 2)
	buf_ch <- "Hello"
	buf_ch <- "World"

	fmt.Println(<-buf_ch)
	fmt.Println(<-buf_ch)
}
