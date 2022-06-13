package main

import (
	"errors"
	"fmt"
)

func greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can not be null")
	}
	return fmt.Sprintf("Hello %s.", name), nil
}

func main() {
	userName := ""
	greet, err := greeting(userName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(greet)
}
