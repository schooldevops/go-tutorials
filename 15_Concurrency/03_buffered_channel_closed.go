package main

import (
	"fmt"
)

func writeCh(buf_ch chan int) {
	for i := 0; i < 5; i++ {
		buf_ch <- i
		fmt.Println("Write data to channel", i)
	}
	close(buf_ch)
}

func main() {
	buf_ch := make(chan int, 10)

	go writeCh(buf_ch)
	for {
		value, ok := <-buf_ch
		if ok == false {
			fmt.Println("Channel is closed")
			break
		}
		fmt.Println("Read value from buf", value)
	}
}
