package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 1

	//constant time
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]

	fmt.Println(s)
}
