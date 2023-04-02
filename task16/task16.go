package main

import "fmt"

func qwik_sort(lst []int, left int, right int) []int {

	//Создаем копии пришедших переменных, с которыми будем манипулировать в дальнейшем.
	l := left
	r := right

	//Вычисляем 'центр', на который будем опираться. Берем значение ~центральной ячейки массива.
	center := lst[(left+right)/2]

	fmt.Println(l, r, (left+right)/2)

	//Цикл, начинающий саму сортировку
	for l <= r {

		//Ищем значения больше 'центра'
		for lst[r] > center {
			r--
		}

		//Ищем значения меньше 'центра'
		for lst[l] < center {
			l++
		}

		//После прохода циклов проверяем счетчики циклов
		if l <= r {
			//И если условие true, то меняем ячейки друг с другом.
			lst[r], lst[l] = lst[l], lst[r]
			l++
			r--
		}
	}

	if r > left {
		qwik_sort(lst, left, r)
	}

	if l > right {
		qwik_sort(lst, l, right)
	}

	return lst
}

func main() {
	lst := []int{5, 4, 1, 2, 0, 123, 1234, 32, 12, 2345, 99}
	lst = qwik_sort(lst, 0, len(lst)-1)
	fmt.Println(lst)
}
