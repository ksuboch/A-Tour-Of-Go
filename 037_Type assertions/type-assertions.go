package main

import "fmt"

// t := i.(T) этот оператор утверждает, что
// интерфейс i содержит тип T и присваивает
// значение T переменной t

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// ok содержит результат преобразования
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
