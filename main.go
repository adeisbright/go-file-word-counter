package main

import (
	"fmt"

	"github.com/adeisbright/go-file-word-counter/reader"
)

func main() {
	fmt.Println("Welcome to File Word Counter")
	wordCount := reader.Read("Hello Bayo", "Bayo")
	fmt.Println("Bayo occurs ", wordCount)
}
