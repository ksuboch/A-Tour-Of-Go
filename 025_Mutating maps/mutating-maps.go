package main

import "fmt"

func main() {
	m := make(map[string]int)

	// добавление элемента
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// изменение элемента
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// удаление элемента
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// если ключ присутствует - в ok
	// будет true иначе - false
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
