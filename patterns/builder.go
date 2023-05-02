package main

import "fmt"

//Паттерн Builder (Строитель) — это порождающий паттерн проектирования, который предоставляет гибкое
//и инкрементальное создание сложных объектов.Он разделяет процесс создания объекта от его представления.
//Builder особенно полезен, когда у объекта много параметров, и большинство из них являются необязательными.
//Паттерн Builder используется в следующих случаях:
//
//Когда объект имеет множество параметров, и большинство из них являются необязательными.
//Builder позволяет упростить создание объекта и избежать большого количества конструкторов с разными аргументами.
//
//Когда процесс создания объекта состоит из нескольких шагов и разных представлений.
//Builder предоставляет гибкость в создании объектов и позволяет разделить процесс создания от представления объекта.
//
//Когда необходимо инкапсулировать сложную логику создания объекта.
//Builder скрывает детали создания объекта и предоставляет удобный интерфейс для его конфигу
//
//Когда объекты могут иметь разные вариации или представления, но процесс создания этих объектов состоит из одних и тех же шагов.
//В этом случае, можно использовать несколько конкретных строителей, реализующих один и тот же интерфейс Builder,
//для создания разных вариаций объекта.
//
//Когда объекты должны создаваться в разное время или в разных контекстах.
//Builder позволяет отделить процесс создания объекта от его использования, что может быть полезным, когда
//объекты создаются по запросу или в разных средах выполнения.

// Product - это сложный объект, который мы хотим построить.
type Product struct {
	field1 string
	field2 int
	field3 bool
}

// Builder - это интерфейс, описывающий методы для построения объекта Product.
type Builder interface {
	SetField1(string) Builder
	SetField2(int) Builder
	SetField3(bool) Builder
	Build() *Product
}

// ConcreteBuilder - это реализация интерфейса Builder.
type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) SetField1(value string) Builder {
	b.product.field1 = value
	return b
}

func (b *ConcreteBuilder) SetField2(value int) Builder {
	b.product.field2 = value
	return b
}

func (b *ConcreteBuilder) SetField3(value bool) Builder {
	b.product.field3 = value
	return b
}

func (b *ConcreteBuilder) Build() *Product {
	return &b.product
}

func main() {
	builder := &ConcreteBuilder{}
	product := builder.SetField1("Hello").
		SetField2(42).
		SetField3(true).
		Build()

	fmt.Printf("Product: %+v\n", product)
}
