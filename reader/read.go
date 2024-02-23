package reader

import "strings"

func Read(sentence string, text string) int {
	wordCounter := 0
	words := strings.Fields(sentence)
	for _, word := range words {
		if word == text {
			wordCounter = wordCounter + 1
		}
	}
	return wordCounter
}
