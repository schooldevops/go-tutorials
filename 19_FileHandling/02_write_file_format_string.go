package main

import (
	"fmt"
	"os"
)

func main() {
	fh, err := os.Create("string02.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(fh, "Hello %s\n", "Kido")
	fmt.Fprintf(fh, "How old are you? %d, %s?\n", 29, "Kido")

	err = fh.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
