package main

import "fmt"

func main() {
	// Создаем слайс
	celsius := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// создаем мапу для хранения упорядоченных значений
	mapa := make(map[int][]float64)

	for _, value := range celsius {
		//-25 - (- 5)
		i := int(value) - (int(value) % 10)
		//добавляем элемент в слайс
		mapa[i] = append(mapa[i], value)
	}

	fmt.Println(mapa)
}
