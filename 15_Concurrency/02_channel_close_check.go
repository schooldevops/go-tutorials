package main

import "fmt"

func sendSeries(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}

	close(ch)
}

func main() {
	ch := make(chan int)

	go sendSeries(ch)
	for {
		value, ok := <-ch
		if ok == false {
			fmt.Println("Channel already closed")
			break
		}

		fmt.Println("read value from ch:", value)
	}
}
