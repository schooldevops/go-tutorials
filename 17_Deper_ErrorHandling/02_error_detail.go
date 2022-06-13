package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Println("Fail to open file", pathErr.Path)
			return
		}
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Success Open File", f.Name())
}
