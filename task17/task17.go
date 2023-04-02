package main

import (
	"fmt"
)

func binary_search(array []int, value int) int {
	low := 0
	high := len(array) - 1

	for low < high {
		mid := (low + high) / 2
		guess := array[mid]
		if guess == value {
			return mid
		} else if guess > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if array[low] == value {
		return low
	}
	return -1
}
func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	predict := binary_search(array, 11)
	fmt.Printf("%v", predict)
}
