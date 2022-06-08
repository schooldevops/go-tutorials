package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // 데드락이 발생되는 지점
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
