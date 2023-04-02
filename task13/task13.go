package main

import "fmt"

func main() {
	//Поменять местами два числа без создания временной переменной
	//a, b = b, a - не подходит, создается временный кортеж - https://www.bestprog.net/ru/2019/02/26/python-assignment-operator-assignment-forms-examples-positional-assignment-of-tuples-lists-ru/
	a := 1
	b := 2
	fmt.Printf("initial values a:%v, b:%v\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("after swap values a:%v, b:%v\n", a, b)

	//Для нецелочисленных:
	c := 1.5
	d := 5.7
	fmt.Printf("initial values c:%v, d:%v\n", c, d)
	c = c * d
	d = c / d
	c = c / d
	fmt.Printf("after swap values c:%v, d:%v\n", c, d)

	x := false
	y := true
	//XOR
	fmt.Printf("initial values x:%v, y:%v\n", x, y)
	x = (x || y) && !(x && y)
	y = (x || y) && !(x && y)
	x = (x || y) && !(x && y)
	fmt.Printf("after swap values x:%v, y:%v\n", x, y)
}
