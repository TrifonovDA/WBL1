package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.Pow(3, 40)
	b := math.Pow(2, 50)
	fmt.Printf("Вычитание: %v\n", a-b)
	fmt.Printf("Сумма: %v\n", a+b)
	fmt.Printf("Умножение: %v\n", a*b)
	fmt.Printf("Деление: %v\n", a/b)
}
