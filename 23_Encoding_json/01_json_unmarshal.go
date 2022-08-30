package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var jsonBlob = []byte(
		`[
			{"Name": "Platypus", "Order": "Monotremata"},
			{"Name": "Quoll", "Order": "Dasyuromorphia"}
		]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", animals)
}
