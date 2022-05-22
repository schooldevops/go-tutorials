package main

import (
	"fmt"
)

func main() {
	const (
		PI         = 3.1415
		BASE_AGE   = 10
		NAME       = "KIDO"
		FORAMT_STR = "Hello %s, PI is %f"
	)
	fmt.Println(PI)
	fmt.Println(BASE_AGE)
	fmt.Println(fmt.Sprintf(FORAMT_STR, NAME, PI))
}
