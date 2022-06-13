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

	reader := bufio.NewScanner(fh)
	for reader.Scan() {
		fmt.Println(reader.Text())
	}

	err = reader.Err()
	if err != nil {
		log.Fatal("read)")
	}
}
