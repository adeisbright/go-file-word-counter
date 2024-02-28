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
	//Getting the Current Directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	fmt.Println("Working Directory is", wd)

	//Reading From a File at Once
	reader.ReadFromFile("example.txt")
	//Create a file
	// f, err := os.Create("hello.txt")
	// if err != nil {
	// 	fmt.Println("Problem Creating File: ", err)
	// 	return
	// }
	// defer f.Close()
	// f.WriteString("Example of How to Write to a File")
	// fmt.Println("Program Successfully  Completed")
	reader.WriteToFile("hello.txt", "John is coming\n")
	reader.WriteToFile("hello.txt", "John is going\n")
}
