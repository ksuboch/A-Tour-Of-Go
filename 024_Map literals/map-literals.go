package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.58433, -74.39967,
	},
	// для типов верхнего уровня можно не указывать имя
	"Google": {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
