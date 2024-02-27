package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/adeisbright/go-file-word-counter/reader"
)

func main() {
	fmt.Println("Welcome to File Word Counter")
	sentence := "The quick brown fox jumps over the lazy dog"
	words := strings.Fields(sentence)
	word := "dog"
	wordCount, err := reader.Read(words, word)
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
	fmt.Printf("The word '%s' appears %d times in the sentence.\n", word, wordCount)
	args := os.Args
	fileName, err := reader.ReadCommandLine(args)
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	fmt.Println("The filename is", fileName)

}
