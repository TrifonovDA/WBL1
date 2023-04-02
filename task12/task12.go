package main

import "fmt"

func to_set(slice []string) (set []string) {
	slice_to_set := make(map[string]bool)
	for _, val := range slice {
		slice_to_set[val] = true
	} // вообще это можно считать сетом

	set = make([]string, 0, len(slice_to_set))
	for k := range slice_to_set {
		set = append(set, k)
	}
	return set
}

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(to_set(slice))
}
