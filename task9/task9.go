package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var x []int
	for i := 0; i < 100; i++ {
		x = append(x, i)
	}

	go func() {
		for {
			select {
			case val, ok := <-ch1:
				if !ok {
					continue
				}
				ch2 <- val * val
			}
		}
	}()

	go func() {
		for {
			select {
			case val, ok := <-ch2:
				if !ok {
					continue
				}
				fmt.Println(val)
			}
		}
	}()

	for _, val := range (x) {
		ch1 <- val
	}
}
