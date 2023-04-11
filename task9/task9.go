package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)

	var x []int
	for i := 1; i < 100; i++ {
		x = append(x, i)
	}

	go func() {
		for {
			select {
			case val, ok := <-ch1:
				if !ok {
					continue
				}
				ch2 <- val + 2
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
				wg.Done()
			}
		}
	}()

	for _, val := range x {
		wg.Add(1)
		ch1 <- val
	}
	wg.Wait()
	//time.Sleep(1 * time.Second)
}
