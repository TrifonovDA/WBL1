package main

import "fmt"

func main() {
	var initial_word string
	initial_word = "главрыба"
	runes := []rune(initial_word)
	var new_word = make([]rune, len(runes))

	for j := len(runes) - 1; j > -1; j-- {
		new_word = append(new_word, runes[j])
	}
	fmt.Println(string(new_word))
}
