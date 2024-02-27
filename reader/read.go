package reader

import (
	"errors"
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
