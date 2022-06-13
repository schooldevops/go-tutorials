package main

import (
	"fmt"
	"os"
)

func main() {
	fh, err := os.Create("byte01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	greeting := "Hello World"
	greeting_byte := []byte(greeting)

	// greeting_byte := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}

	writeByte, err := fh.Write(greeting_byte)
	if err != nil {
		fmt.Println(err)
		fh.Close()
		return
	}

	fmt.Println("write byte is ", writeByte)
	err = fh.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
