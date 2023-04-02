package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

func main() {
	var wg sync.WaitGroup
	var sum atomic.Uint32
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			sum.Swap(sum.Add(1))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}
