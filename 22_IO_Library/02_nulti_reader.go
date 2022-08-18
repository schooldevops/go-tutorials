package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("First Data.....")
	r2 := strings.NewReader("Second Data.....")
	r3 := strings.NewReader("Third Data.....")

	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
