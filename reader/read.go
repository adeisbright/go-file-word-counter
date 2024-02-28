package reader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Read(sentence []string, text string) (int, error) {
	if text == "" || len(sentence) == 0 {
		return 0, errors.New("sententce or words cannot be empty")
	}
	wordCounter := 0
	for _, word := range sentence {
		if word == text {
			wordCounter = wordCounter + 1
		}
	}
	return wordCounter, nil
}

func ReadCommandLine(args []string) (string, error) {
	if len(args) < 2 {
		return "a", errors.New("pass the name of the file in the cli")
	}
	fileName := args[1]
	return fileName, nil
}

func WriteToFile(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//Wrap the file with a bufio for buffered IO
	writer := bufio.NewWriter(file)
	//Write the content to the file
	_, err = writer.WriteString(content)
	if err != nil {
		fmt.Println(err)
	}
	//Flush to the buffer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
	}
}

func ReadFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Printf("%s has been created", fileName)
}

func RemoveFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s has been removed", fileName)
}

func RenameFile(oldName string, newName string) {
	err := os.Rename(oldName, newName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s has been removed to %s", oldName, newName)
}
