package main

import (
	"fmt"
	"sync"
	"time"
)

func process(number int, wg *sync.WaitGroup) {
	fmt.Println("Start process number", number)
	time.Sleep(2 * time.Second)
	fmt.Println("End process number", number)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routine are done.")
}
