package main

import (
	"fmt"
	"time"
)

func main() {
	d := 10 * time.Second
	mySleep(d)
	fmt.Println("sleep is over")
}

func mySleep(d time.Duration) {
	<-time.After(d)
}
