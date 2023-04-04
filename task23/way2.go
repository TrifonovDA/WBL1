package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 1

	//linear time due to copy
	copy(s[i:], s[i+1:])
	s = s[:len(s)-1]

	fmt.Println(s)
}
