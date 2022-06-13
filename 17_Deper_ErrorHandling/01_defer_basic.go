package main

import "fmt"

func defer_statement() {
	fmt.Println("-------- End by Defer ---------")
}

func process() {
	defer defer_statement()

	for i := 0; i < 10; i++ {
		fmt.Println("Print value :", i)
	}
}

func main() {
	process()
}
