package main

import "fmt"

// Паттерн "Фабрика" (Factory) — это порождающий паттерн проектирования, который предоставляет интерфейс для
//создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.
//Это обеспечивает гибкость в создании объектов и инкапсуляцию логики создания объектов.

//Паттерн "Фабрика" используется в следующих случаях:
//
//Когда объекты могут быть созданы с различными вариациями или представлениями, но процесс создания этих объектов состоит из одних и тех же шагов. В этом случае, фабрика может быть использована для создания разных вариаций объекта, инкапсулируя логику создания.
//
//Когда создание объекта требует сложной логики или зависит от других объектов или параметров. Фабрика инкапсулирует эту логику и предоставляет удобный интерфейс для создания объектов.
//
//Когда необходимо отделить процесс создания объекта от его использования.

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

type AnimalFactory interface {
	CreateAnimal() Animal
}

type DogFactory struct{}

func (df DogFactory) CreateAnimal() Animal {
	return Dog{}
}

type CatFactory struct{}

func (cf CatFactory) CreateAnimal() Animal {
	return Cat{}
}

func main() {
	var dogFactory AnimalFactory = DogFactory{}
	var catFactory AnimalFactory = CatFactory{}

	dog := dogFactory.CreateAnimal()
	cat := catFactory.CreateAnimal()

	fmt.Println(dog.Speak())
	fmt.Println(cat.Speak())
}
