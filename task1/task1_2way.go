package main

import "fmt"

type human_interface2 interface {
	Get_name2()
}

// Задаем структуру Human
type human2 struct {
	name    string
	surname string
}

// Назначаем структуре метод. На структуру не ссылаемся, так как нам не нужно ее менять.
func (h human2) Get_name2() {
	fmt.Println("Hi! My name is", h.name, h.surname)
}

func Say_trash2(h human_interface2) {
	fmt.Println("$%@#@@#$%!")
}

// Задаем дочернюю структуру Action
type Action2 struct {
	human2 //встраиваем родительскую структуру в дочернюю.
	// Важно встраивать именно так, а не "human: Human", чтобы был доступ к методам родителя
	process string
}

func main() {
	parrent := human2{name: "Masha", surname: "Ivanova"} // Инициализируем родительскую структуру
	child := Action2{
		human2:  parrent,
		process: "sleeping",
	} //задаем детскую структуру, в которой указана принадлежность к родителю
	fmt.Println("Child get name of parrent essense: ")
	child.Get_name2() //Вызываем родительский метод.
	fmt.Println("Child's action is", child.process)
	//Proof of problem:
	Say_trash2(child)
	//научили ребенка ругаться
}

//как работает:
//функция Say_trash принимает на вход интерфейс. Так как и детская сущность перенимает методы родительской сущности, то она тоже удовлетворяет интерфейсу.
