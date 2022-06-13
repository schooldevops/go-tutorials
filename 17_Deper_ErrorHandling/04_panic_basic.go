package main

import "fmt"

func greeting(username *string) {
	if username == nil {
		panic("username cannot be nil")
	}

	fmt.Println(*username, "Nice meet U")
}

func main() {
	name := "Kido"
	greeting(&name)
	greeting(nil)

	fmt.Println("Program -------- exit ----------")
}
