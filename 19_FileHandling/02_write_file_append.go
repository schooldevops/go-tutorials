package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fh, err := os.OpenFile("append.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	dt := time.Now()
	fmt.Fprintf(fh, "[%s] Hello %s\n", dt, "Kido")
	fmt.Fprintf(fh, "[%s] How old are you? %d, %s?\n", dt, 29, "Kido")

	err = fh.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
