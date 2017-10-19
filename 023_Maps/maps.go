package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// карты ассоциирую ключ и значение
var m map[string]Vertex

func main() {
	// нулевое значение карты nil
	// нулевая карта не содержит ключи
	// ключи в неё не могут быть добавлены
	// карту нужно инициализировать функцией make()
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
