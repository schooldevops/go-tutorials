package main

import (
	"fmt"
	"unicode"
)

type CommonInterface interface {
	Length() int
	Capital() string
}

type String string

func (s String) Length() int {
	return len(s)
}

func (s String) Capital() string {
	var resultStr []rune
	for i, rune := range s {
		if i == 0 {
			resultStr = append(resultStr, unicode.ToUpper(rune))
		} else {
			resultStr = append(resultStr, rune)
		}
	}
	return string(resultStr)
}

func main() {
	fruit := String("banana")
	var v CommonInterface
	v = fruit

	fmt.Println("Capital String:", v.Capital())
}
