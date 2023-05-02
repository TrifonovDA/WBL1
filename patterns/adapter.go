package main

import "fmt"

// Паттерн "Адаптер" (Adapter) - это структурный паттерн проектирования, который позволяет объектам с
//несовместимыми интерфейсами работать вместе.
//
//Паттерн "Адаптер" используется в следующих случаях:
//
//Когда необходимо использовать существующий класс, но его интерфейс не соответствует требуемому.
//Когда необходимо интегрировать несколько классов или модулей с разными интерфейсами, чтобы они могли работать вместе.
//Когда необходимо адаптировать интерфейсы классов для работы с внешними системами или сторонними библиотеками.

type Target interface {
	Request() string
}

type Adaptee struct {
	SpecificRequest string
}

func (a Adaptee) SpecificRequestFunction() string {
	return a.SpecificRequest
}

type Adapter struct {
	adaptee Adaptee
}

func (a Adapter) Request() string {
	return a.adaptee.SpecificRequestFunction()
}

func main() {
	adaptee := Adaptee{SpecificRequest: "Adaptee SpecificRequest"}
	adapter := Adapter{adaptee: adaptee}
	fmt.Println(adapter.Request())
}
