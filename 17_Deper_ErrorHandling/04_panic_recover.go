package main

import "fmt"

func recoverUserName() {
	if r := recover(); r != nil {
		fmt.Println("Recover :", r)
	}
}

func greeting(username *string) {
	defer recoverUserName()
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
