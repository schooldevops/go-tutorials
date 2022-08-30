package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Id   int
	Name string
	Age  int
	Tall int
}

type userInfo2 struct {
	Id   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age"`
	Tall int    `json:"tall"`
}

func main() {

	user01 := &userInfo{
		Id:   1,
		Name: "Kido",
		Age:  30,
		Tall: 180,
	}
	user01_m, _ := json.Marshal(user01)
	fmt.Println(string(user01_m))

	user02 := &userInfo2{
		Id:   2,
		Name: "Kido2",
		Age:  33,
		Tall: 177,
	}
	user02_m, _ := json.Marshal(user02)
	fmt.Println(string(user02_m))

	user03 := []byte(`{"id":3,"name":"kido3", "age": 40, "tall": 181}`)

	var user03_m map[string]interface{}

	if err := json.Unmarshal(user03, &user03_m); err != nil {
		panic(err)
	}
	fmt.Println(user03_m)

	num := user03_m["age"].(float64)
	fmt.Println(num)

	str := `{"id":4,"name":"kido4", "age": 46, "tall": 176}`
	res := userInfo2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Name)
}
