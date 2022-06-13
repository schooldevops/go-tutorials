package main

import (
	"fmt"
	"os"
)

func main() {
	fh, err := os.Create("string01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	writeByte, err := fh.WriteString("Hello World")
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
