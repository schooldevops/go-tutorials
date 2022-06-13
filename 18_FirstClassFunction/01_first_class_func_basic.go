package main

import "fmt"

func main() {
	firstClass := func() {
		fmt.Println("Print 01")
	}

	firstClass()
	fmt.Println("firstClass: ", firstClass)
	fmt.Printf("firstClass with %%T, %T \n", firstClass)

}
