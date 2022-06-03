package main

import "fmt"

func changeString(s string) string {
	// s[0] = 'A' <-- 컴파일 오류 발생
	return s
}

func changeRune(s []rune) string {
	s[0] = 'A'
	return string(s)
}

func main() {
	greeting := "Hello World"
	fmt.Println(changeRune([]rune(greeting)))
	fmt.Println(changeString(greeting))
}
