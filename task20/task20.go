package main

import (
	"fmt"
	"strings"
)

func main() {
	example := "snow dog sun"

	fmt.Println(reverseWords(example))
}

func reverseWords(s string) string {
	wordsList := strings.Split(s, " ")
	wordsListLength := len(wordsList)

	for i := 0; i < wordsListLength/2; i++ {
		wordsList[i], wordsList[wordsListLength-i-1] = wordsList[wordsListLength-i-1], wordsList[i] //меняем слова местами
	}

	return strings.Join(wordsList, " ")
}
