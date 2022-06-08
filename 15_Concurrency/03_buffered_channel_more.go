package main

import (
	"fmt"
	"time"
)

func writeBuffer(buf_ch chan int) {
	for i := 0; i < 5; i++ {
		buf_ch <- i
		fmt.Println("write to buffer", i)
	}

	close(buf_ch)
}

func main() {
	// 버퍼 2개를 가진 정수형 타입 채널을 생성한다.
	buf_ch := make(chan int, 2)

	go writeBuffer(buf_ch)
	time.Sleep(2 * time.Second)
	for v := range buf_ch {
		fmt.Println("Read from buffer ch", v)
		time.Sleep(2 * time.Second)
	}
}
