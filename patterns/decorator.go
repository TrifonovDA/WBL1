package main

import "fmt"

//Паттерн "Декоратор" (Decorator) - это структурный паттерн проектирования, который позволяет динамически добавлять
//новые функциональности к объектам, не изменяя их код.

//Паттерн "Декоратор" используется в следующих случаях:
//
//Когда необходимо динамически добавлять функциональность к объектам без изменения их кода.
//Когда нельзя наследовать классы или интерфейсы, чтобы добавить функциональность, например,
//когда классы или интерфейсы сделаны финальными.
//Когда расширение функциональности с помощью наследования приводит к созданию большого числа подклассов и
//усложняет архитектуру системы.

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (c ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

type Decorator interface {
	Component
}

type ConcreteDecoratorA struct {
	component Component
}

func (d ConcreteDecoratorA) Operation() string {
	return "ConcreteDecoratorA(" + d.component.Operation() + ")"
}

type ConcreteDecoratorB struct {
	component Component
}

func (d ConcreteDecoratorB) Operation() string {
	return "ConcreteDecoratorB(" + d.component.Operation() + ")"
}

func main() {
	component := ConcreteComponent{}
	decoratorA := ConcreteDecoratorA{component: component}
	decoratorB := ConcreteDecoratorB{component: decoratorA}
	fmt.Println(decoratorB.Operation())
}
