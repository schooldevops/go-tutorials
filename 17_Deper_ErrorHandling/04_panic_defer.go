package main

import "fmt"

func greeting(username *string) {
	defer fmt.Println("----- greeting end ----")
	if username == nil {
		panic("username cannot be nil")
	}

	fmt.Println(*username, "Nice meet U")
}

func main() {
	defer fmt.Println("---- end of main ----")
	greeting(nil)

	fmt.Println("Program -------- exit ----------")
}
