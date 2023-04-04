package main

import (
	"fmt"
	"sync"
)

// исправил код. wg надо было передавать по ссылке или вообще убрать его из параметров функции.
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
