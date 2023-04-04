package main

import (
	"WBL1/task24/point"
	"fmt"
)

func main() {
	a := point.NewPoint(1, 2)
	b := point.NewPoint(5, 4)

	fmt.Println(point.FindDistance(a, b))
}
