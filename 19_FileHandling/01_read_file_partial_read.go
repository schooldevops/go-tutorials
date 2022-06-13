package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fh, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}

	defer fh.Close()

	reader := bufio.NewReader(fh)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data from file", err)
			break
		}
		fmt.Println(string(buffer[0:n]))
	}
}
