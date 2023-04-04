package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	x := big.NewFloat(math.Pow(3, 40))
	y := big.NewFloat(math.Pow(2, 50))
	var sum big.Float
	fmt.Println(sum.Add(x, y))
	var sub big.Float
	fmt.Println(sub.Sub(x, y))
	var mul big.Float
	fmt.Println(mul.Mul(x, y))
	var div big.Float
	fmt.Println(div.Quo(x, y))
}
