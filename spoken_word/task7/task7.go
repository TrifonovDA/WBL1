package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[0] = 1000
	m[-1] = 124
	m[-100] = 281

	fmt.Println(m)
}
