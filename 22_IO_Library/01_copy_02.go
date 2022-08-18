package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	buff := make([]byte, 1024)

	sourceFile, err := os.Open("source.txt")
	if err != nil {
		fmt.Errorf("Fail to oepn source file/n", err)
	}

	defer sourceFile.Close()

	destFile, err := os.OpenFile("dest.txt", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		fmt.Errorf("Fail to open dest file/n", err)
	}

	defer destFile.Close()

	for {
		readSize, err := sourceFile.Read(buff)
		if err != nil && err != io.EOF {
			fmt.Errorf("Error occur when read file data.", err)
		}

		if readSize == 0 {
			fmt.Println("size is 0")
			break
		}

		writeSize, err := destFile.Write(buff[:readSize])
		if err != nil {
			fmt.Errorf("Error occur write data to dest file", err)
		}
		fmt.Printf("WriteFile Size is %d \n", writeSize)
	}

	fmt.Println("Read File and Copy File to DestFile was done.")

}
