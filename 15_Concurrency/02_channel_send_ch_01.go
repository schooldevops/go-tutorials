package main

import "fmt"

func sendData(send_ch chan<- int) {
	send_ch <- 10
}

func main() {
	send_ch := make(chan<- int)
	go sendData(send_ch)

	// 쓰기전용 채널이다. 아래 코드는 오류가 발생한다.
	fmt.Println(<-send_ch)
}
