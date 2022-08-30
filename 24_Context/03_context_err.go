package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time out occur...")
			err := ctx.Err()
			fmt.Println("Err: ", err)
			return
		default:
			fmt.Println("do something ..... ")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start Context with Err...")
	defer fmt.Println("End of example.")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doSomething(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Process have exceeded dead line.")
	}

	time.Sleep(2 * time.Second)
}
