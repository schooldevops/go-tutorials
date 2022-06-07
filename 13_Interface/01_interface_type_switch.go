package main

import "fmt"

type Student struct {
	name    string
	korean  int
	english int
	math    int
}

func printValueByType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("Type is string value is %s\n", i.(string))
	case int:
		fmt.Printf("Type is int value is %d\n", i.(int))
	case Student:
		fmt.Printf("Type is Student value is %v\n", i.(Student))
	default:
		fmt.Println("Unkonwn type")
	}
}

func main() {
	printValueByType("Hello")
	printValueByType(123)
	printValueByType(Student{"Kido", 95, 90, 85})
	printValueByType(true)
}
