package main

import (
	"fmt"

	"github.com/schooldevops/go-module-repo/greeting"
)

func main() {
	greet := greeting.Greeting("Schooldevops")
	fmt.Println(greet)
}
