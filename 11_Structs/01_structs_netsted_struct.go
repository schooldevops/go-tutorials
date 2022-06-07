package main

import "fmt"

type Struct struct {
	name      string
	korean    int
	english   int
	math      int
	bodyIndex BodyIndex
}

type BodyIndex struct {
	weight float32
	height float32
}

func main() {
	kido := Struct{
		name:    "Kido",
		korean:  95,
		english: 90,
		math:    80,
		bodyIndex: BodyIndex{
			weight: 80,
			height: 177,
		},
	}

	fmt.Println("Name is", kido.name)
	fmt.Println("Korean is", kido.korean)
	fmt.Println("English is", kido.english)
	fmt.Println("Math is", kido.math)
	fmt.Println("Weight is", kido.bodyIndex.weight)
	fmt.Println("Height is", kido.bodyIndex.height)
}
