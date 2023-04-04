package main

import "fmt"

func main() {
	//ответ: для среза 3, для мапы 2
	a := make(map[int]int)
	a[1] = 2
	fmt.Printf("type: %T    value: %v\n", a, a)

	var b = map[int]int{1: 1, 2: 3}
	fmt.Printf("type: %T    value: %v\n", b, b)

	//создание среза из массива
	var c = [5]int{1, 2, 3, 4, 5}
	var c_slice = c[:]
	fmt.Printf("type: %T    value: %v\n", c_slice, c_slice)

}
