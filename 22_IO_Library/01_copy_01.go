package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	sourceFile, err := os.Open("source.txt")

	if err != nil {
		fmt.Println("Error occur during read source file")
	}

	defer sourceFile.Close()

	destFile, err := os.OpenFile("dest.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error make dest file")
	}

	defer destFile.Close()

	readByteLen, err := io.Copy(io.Writer(destFile), io.Reader(sourceFile))

	fmt.Println("Read File and Copy File to DestFile. readByte: ", readByteLen)

}
