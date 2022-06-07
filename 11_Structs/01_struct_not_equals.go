package main

import (
	"fmt"
)

type userInfos struct {
	data map[int]int
}

func main() {
	userInfos1 := userInfos{
		data: map[int]int{
			0: 90,
		}}
	userInfos2 := userInfos{
		data: map[int]int{
			0: 90,
		}}

	fmt.Println("userInfos1", userInfos1)
	fmt.Println("userInfos2", userInfos2)
	// 아래 내용은 컴파일 오륙를 일으킨다.
	// if userInfos1 == userInfos2 {
	// 	fmt.Println("userInfos1 and userInfos2 are equal")
	// }
}
