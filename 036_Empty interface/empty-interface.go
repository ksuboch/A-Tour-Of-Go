package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	// interface{} используется в коде
	// для обработки значений неизвестного типа
	// например fmt.Print принимает любое
	// количество аргументов типа interface{}
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}
