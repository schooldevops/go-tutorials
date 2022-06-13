package main

import (
	"error_sample_complex/operror"
	"fmt"
)

func halfOfReaOfRect(width float64, height float64, dividen float64) (float64, error) {
	if width <= 0 || height <= 0 {
		return 0, operror.New("Not valid width or height", width, height, dividen)
	}
	if dividen <= 0 {
		return 0, operror.New("Not valid dividen", width, height, dividen)
	}
	return (width * height) / dividen, nil
}

func main() {
	value, err := halfOfReaOfRect(10, 10, 0)
	if err != nil {
		if err, ok := err.(*operror.OperationError); ok {
			if err.IsZeroDividen() {
				fmt.Println("error: zerodividen")
			}

			if err.IsNotValidValue() {
				fmt.Println("error: not vaild value")
			}
		}
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}
