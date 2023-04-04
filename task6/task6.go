package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("sec.")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	// …
	time.Sleep(5 * time.Second)
	close(quit)
}
