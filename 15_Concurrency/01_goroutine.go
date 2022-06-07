package main

import "fmt"

func greeting() {
	fmt.Println("Hello World. I'm go routine.")
}

func main() {
	go greeting()
	fmt.Println("I'm a main function")
}
