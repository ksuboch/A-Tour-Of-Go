package main

import "fmt"

// Vertex - struct (коллекция полей)
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	// доступ к полям структуры производится с помощью точки
	v.X = 4
	fmt.Println(v.X)

	// доступ к полям может быть положен через указатель
	p := &v
	// можно не делать (*p).X - явного разыменовывания
	p.X = 1e9
	fmt.Println(v)
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}  // Y:0
		v3 = Vertex{}      // X:0 и Y:0
		d  = &Vertex{1, 2} // d имеет тип *Vertex - указатель на структуру
	)
	fmt.Println(v1, d, v2, v3)
}
