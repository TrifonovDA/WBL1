package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan int)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Закрываем канал, завершаем работу горутины.")
				return
			default:
				fmt.Println("sec.")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	quit <- 1
}
