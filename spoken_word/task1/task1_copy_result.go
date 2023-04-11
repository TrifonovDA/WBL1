package main

import (
	"fmt"
)

// https://golangify.com/concatenation - по бенчмарку этот метод быстрее.
func main() {

	strings := []string{"This ", "is ", "even ",
		"more ", "performant "}

	bs := make([]byte, 100) //Проблемное место
	bl := 0

	for _, val := range strings {
		bl += copy(bs[bl:], []byte(val))
	}

	fmt.Println(string(bs[:]))

}
