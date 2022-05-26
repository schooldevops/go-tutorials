package main

import (
	"fmt"
)

func main() {

	i := 0

point2:
	for {
	point:
		for {
			i++
			if i > 10 {
				fmt.Println("i between 5 and 10. breaking to label point2")
				break point2
			} else if i > 5 {
				fmt.Printf("i is %d. breaking to label point\n", i)
				break point
			}
			fmt.Println("Current Value", i)
		}
	}
}
