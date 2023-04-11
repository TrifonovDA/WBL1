package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Закрываем канал, завершаем работу горутины.")
				wg.Done()
				return
			default:
				fmt.Println("sec.")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()

}
