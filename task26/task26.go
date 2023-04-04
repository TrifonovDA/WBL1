package main

import (
	"fmt"
	"strings"
)

func main() {
	a := `abCdefA`

	res := checkCharacters(a)

	fmt.Println(res)
}

func checkCharacters(s string) bool {
	visited := make(map[rune]bool, 0)

	s = strings.ToLower(s)

	sl := []rune(s)

	for _, v := range sl {
		_, b := visited[v]
		if b {
			return false
		}
		visited[v] = true
	}
	return true
}
