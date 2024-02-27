package main

import (
	"bufio"
	"fmt"
	"log"
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
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Problem Opening File: ", err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		processData(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	fmt.Println("Program Successfully  Completed")
}

func processData(line string) {
	fmt.Println(line)
}
