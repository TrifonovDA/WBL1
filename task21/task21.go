package main

import "fmt"

//Паттерн Adapter относится к структурным паттернам уровня класса.
//
//Часто в новом проекте разработчики хотят повторно использовать уже существующий код.
//Например, имеющиеся классы могут обладать нужной функциональностью и иметь при этом несовместимые интерфейсы. В таких случаях следует использовать паттерн Adapter.
//
//Смысл работы этого паттерна в том, что если у вас есть класс и его интерфейс не совместим с кодом вашей системы, то что бы разрешить этот конфликт,
//мы не изменяем код этого класса, а пишем для него адаптер. Другими словами Adapter адаптирует существующий код к требуемому интерфейсу (является переходником).
//
//Требуется для реализации:
//
//Интерфейс Target, описывающий целевой интерфейс (тот интерфейс с которым наша система хотела бы работать);
//Класс Adaptee, который наша система должна адаптировать под себя;
//Класс Adapter, адаптер реализующий целевой интерфейс.

type computer interface {
	connectUSB()
}

type newLaptop struct {
}

func (o *newLaptop) connectUSBTypeC() {
	fmt.Println("connected USB type C...")
}

type oldLaptop struct {
}

func (o *oldLaptop) connectUSB() {
	fmt.Println("connected USB...")
}

type adapter struct {
	laptop newLaptop
}

func (a *adapter) connectUSB() { //адаптер удовлетворяет интерфейсу, но мы заменили ему метод
	a.laptop.connectUSBTypeC()
}

func main() {
	o := oldLaptop{} //создали старый ноутбук
	o.connectUSB()   //вызвали метод connectUSB

	n := newLaptop{} //создали новый ноутбук
	a := adapter{n}  //создали адаптер для нового ноутбука (надели на шнур переходник)
	a.connectUSB()   //вызвали метод connectUSB
}
