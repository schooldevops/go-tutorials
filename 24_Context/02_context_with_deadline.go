package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethingDuringTime(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time out !!!! You Failed.")
			return
		default:
			fmt.Printf("Get key from context... [%v]\n", ctx.Value("my-key"))
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start Context with Timeout...")
	defer fmt.Println("End of example.")

	ctx := context.WithValue(context.Background(), "my-key", "Hello This is key...")
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go doSomethingDuringTime(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Process have exceeded dead line.")
	}

	time.Sleep(2 * time.Second)
}
