package main

//https://habr.com/ru/post/571002/
import "fmt"

// Т.к. функция append не меняет старый, а возвращает новый слайс, приходится
// 1. либо возвращать return'ом новый слайс и с ним работать.
// либо передать указатель на слайс в параметрах функции.
func main() {
	slice := []string{"a", "a"}

	func(slice *[]string) {
		*slice = append(*slice, "a")
		(*slice)[0] = "b"
		(*slice)[1] = "b"
		fmt.Println(slice)
	}(&slice)
	fmt.Println(slice)
}
