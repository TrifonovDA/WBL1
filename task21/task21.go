package main

import "fmt"

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

type person struct {
}

func (p *person) insertUSB(c computer) {
	fmt.Println("inserted USB...")
}

type adapter struct {
	laptop *newLaptop
}

func (a *adapter) connectUSB() {
	a.laptop.connectUSBTypeC()
}

func main() {
	o := &oldLaptop{}
	o.connectUSB()

	n := &newLaptop{}
	a := &adapter{n}
	a.connectUSB()
}
