package main

import (
	"fmt"
	"strconv"
)

//попробовать без пакетов и конвертаций, как второе решение

func change(str string, i uint8, val bool) string {
	array := []rune(str)
	if val == true {
		array[i] = '1'
	} else {
		array[i] = '0'
	}
	fmt.Println(string(array))
	return string(array)
}

func main() {
	var N int64
	var i uint8
	i = 10
	var val bool
	val = true
	N = 5
	str := fmt.Sprintf("%b", N)
	if i > uint8(len(str)) {
		fmt.Println("out of range")
		return
	}
	l, err := strconv.ParseInt(change(str, i, val), 2, 64)
	fmt.Println(l, err)
}
