package main

import (
	"fmt"
	"math/rand"
	"time"
)

func process(ch_01 chan string, ch_02 chan string) {
	time.Sleep(1 * time.Second)
	for i := 0; i < 20; i++ {
		value := rand.Intn(10)
		if value%2 == 0 {
			ch_01 <- fmt.Sprintf("channel 01 %d", value)
		} else {
			ch_02 <- fmt.Sprintf("channel 02 %d", value)
		}
	}
}

func exiter(exit chan bool) {
	time.Sleep(10 * time.Second)
	exit <- true
}

func main() {
	ch_01 := make(chan string)
	ch_02 := make(chan string)
	exit_ch := make(chan bool)

	go process(ch_01, ch_02)
	go exiter(exit_ch)
exit_pnt:
	for {
		time.Sleep(1 * time.Second)
		select {
		case val01 := <-ch_01:
			fmt.Println(val01)
		case val02 := <-ch_02:
			fmt.Println(val02)
		case <-exit_ch:
			fmt.Println("Exit Program")
			break exit_pnt
		default:
			fmt.Println("Default selector")
		}
	}
}
