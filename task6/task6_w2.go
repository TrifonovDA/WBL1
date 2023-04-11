package main

import (
	"fmt"
	"time"
)

// Generator возвращает канал,
// который производит числа 1, 2, 3,…
// Чтобы остановить подлежащую goroutine, закройте канал.
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case _, ok := <-ch:
				if !ok {
					fmt.Println("close")
					return
				}
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	close(number)
	fmt.Println(<-number)
	time.Sleep(time.Second * 3)
}
