package main

import "fmt"

func sendData(send_ch chan<- int) {
	send_ch <- 10
}

func main() {
	send_ch := make(chan int)
	go sendData(send_ch)
	fmt.Println(<-send_ch)
}
