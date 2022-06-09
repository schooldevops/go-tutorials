package main

import (
	"fmt"
	"math/rand"
	"time"
)

func greeting(name string, ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Hello " + name
}

func dice(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- rand.Intn(5) + 1
}

func main() {
	ch_01 := make(chan string)
	ch_02 := make(chan int)

	go greeting("Kido", ch_01)
	go dice(ch_02)

	select {
	case greet := <-ch_01:
		fmt.Println(greet)
	case randVal := <-ch_02:
		fmt.Println("Random Value: ", randVal)
	}
}
