package main

import "fmt"

// Задаем структуру Human
type human struct {
	name    string
	surname string
}

// Назначаем структуре метод. На структуру не ссылаемся, так как нам не нужно ее менять.
func (h human) get_name() {
	fmt.Println("Hi! My name is", h.name, h.surname)
}

func say_trash1(h human) {
	fmt.Println("$%@#@@#$%!")
}

// Задаем дочернюю структуру Action
type action struct {
	human //встраиваем родительскую структуру в дочернюю.
	// Важно встраивать именно так, а не "human: Human", чтобы был доступ к методам родителя
	process string
}

func main() {
	parrent := human{name: "Masha", surname: "Ivanova"} // Инициализируем родительскую структуру
	child := action{
		human:   parrent,
		process: "sleeping",
	} //задаем детскую структуру, в которой указана принадлежность к родителю
	fmt.Println("Child get name of parrent essense: ")
	child.get_name() //Вызываем родительский метод. Mission complete?
	fmt.Println("Child's action is", child.process)
	//Proof of problem:
	//Say_trash1(child)
}

//Проблемы реализации: методы родительской функции работают, а функции родительской сущности не работают.
//Проблема решается с помощью интерфейса.
