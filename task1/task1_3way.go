package main

import "fmt"

// Множественное наследование.
type human_interface13 interface {
	Get_name3()
}
type human_interface23 interface {
	Get_surname3()
}

// Задаем структуру Human
type human13 struct {
	name string
}
type human23 struct {
	surname string
}

// Назначаем структуре метод. На структуру не ссылаемся, так как нам не нужно ее менять.
func (h human13) Get_name3() {
	fmt.Println("Hi! My name is", h.name)
}
func (h human23) Get_surname3() {
	fmt.Println("Hi! My name is", h.surname)
}

func Say_good_name(human human_interface13) {
	fmt.Println("good name")
}
func Say_good_surname(human human_interface23) {
	fmt.Println("good surname")
}

// Задаем дочернюю структуру Action
type action3 struct {
	human13 //встраиваем родительскую структуру в дочернюю.
	human23 //встраиваем родительскую структуру в дочернюю.
	// Важно встраивать именно так, а не "human: Human", чтобы был доступ к методам родителя
	process string
}

func main() {
	parrent1 := human13{name: "Masha"} // Инициализируем родительскую структуру
	parrent2 := human23{surname: "Ivanova"}
	child := action3{
		human13: parrent1,
		human23: parrent2,
		process: "sleeping",
	} //задаем детскую структуру, в которой указана принадлежность к родителям

	child.Get_name3()    //Вызываем родительский метод от первого родителя
	child.Get_surname3() //Вызываем родительский метод от второго родителя
	Say_good_name(child)
	Say_good_surname(child)
}
