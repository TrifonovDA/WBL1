package main

import "fmt"

func intersection(first_slice []int, second_slice []int) (result_slice []int) {
	mapa := make(map[int]bool)

	if len(first_slice) < len(second_slice) {
		for _, val := range first_slice {
			mapa[val] = true
		}
		for _, val := range second_slice {
			counter, ok := mapa[val]
			if ok && counter != false {
				result_slice = append(result_slice, val)
				mapa[val] = false
			}
		}
		return result_slice
	} else {
		for _, val := range second_slice {
			mapa[val] = true
		}
		for _, val := range first_slice {
			counter, ok := mapa[val]
			if ok && counter != false {
				result_slice = append(result_slice, val)
				mapa[val] = false
			}
		}
		return result_slice
	}
}

func main() {
	var first_slice = []int{1, 4, 5, 7, 2, 9}
	var second_slice = []int{3, 10, 1, 2}

	fmt.Println(intersection(first_slice, second_slice))
}
