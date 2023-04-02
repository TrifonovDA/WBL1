package main

import (
	"fmt"
	"io"
	"os"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.

// var justString string --
// К глобальным переменным можно обращаться из любой функции, поэтому их изменение может быть неожиданным.
// Такое поведение может привести к ошибкам, которые очень сложно отлаживать.
// На самом деле, глобальные переменные допустимы, но их использование должно быть оправдано.
// Например, в некоторых случаях они могут быть полезны для конфигурации приложения, но в большинстве случаев лучше использовать локальные переменные.
// Локальные переменные создаются на стеке, а глобальные - на куче.
// Работа со стеком быстрее, чем с кучей, поэтому лучше использовать локальные переменные, если это возможно.
//
//	func someFunc() {
//		v := createHugeString(1 << 10)
//
//		justString = v[:100]
//	}
//
//	func main() {
//		someFunc()
//	}
func createHugeString(uint uint64, writer io.Writer) (string, error) {
	var str []byte
	for i := uint64(0); i < uint; i++ {
		str = append(str, []byte("привет")...)
		_, err := writer.Write([]byte("привет"))
		if err != nil {
			return "", err
		}
	}
	return string(str), nil
}

func someFunc([]byte) []byte {
	f, err := os.Create("huge_string.txt") // для работы с очень большими строками, которые могут не поместиться в памяти,
	// стоит их сохранять на железо, например, в файл
	if err != nil {
		panic(err)
	}
	defer f.Close()
	str, err := createHugeString(1<<20, f)
	if err != nil {
		fmt.Println("Err by printing.", err)
	}
	return []byte(str)
}

func main() {
	var justString []byte
	//строки в Go иммутабельны, то есть когда мы их изменяем, фактически создаём новую строку. Поэтому, если мы хотим изменить строку, то лучше использовать []byte:
	//Локальные переменные создаются на стеке, а глобальные - на куче.
	//Работа со стеком быстрее, чем с кучей, поэтому лучше использовать локальные переменные, если это возможно.
	justString = someFunc(justString)[:100]
	fmt.Println(string(justString))

}
