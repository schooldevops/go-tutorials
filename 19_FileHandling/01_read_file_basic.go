package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	all_data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}
	fmt.Println("------------ Read file ---------")
	fmt.Println(string(all_data))
}
