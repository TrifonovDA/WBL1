package main

import (
	"fmt"
	"reflect"
)

func GetType(a interface{}) string {
	ans := ""
	switch a.(type) {
	case int:
		ans = "int"
	case string:
		ans = "string"
	case bool:
		ans = "bool"
	case chan interface{}:
		ans = "chan interface{}"
	default:
		ans = reflect.TypeOf(a).String()
	}

	return ans
}
func main() {
	//Разработать программу, которая в рантайме способна определить тип переменной:
	//int, string, bool, channel из переменной типа interface{}
	a := make(chan interface{})
	//fmt.Printf("%p", a)
	fmt.Println(GetType(a))

}
