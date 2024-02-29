package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/adeisbright/go-file-word-counter/randomness"
	"github.com/adeisbright/go-file-word-counter/reader"
)

func main() {
	fmt.Println("Welcome to File Word Counter")
	args := os.Args
	fileName, err := reader.ReadCommandLine(args)
	if err != nil {
		fmt.Println("Error Reading Command Line Arg: ", err)
	}
	sentence, err := reader.ReadFromFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	words := strings.Fields(sentence)
	word := "John"
	wordCount, err := reader.Read(words, word)
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
	fmt.Printf("The word '%s' appears %d times in the sentence.\n", word, wordCount)

	randomness.GenerateNumber()
}
