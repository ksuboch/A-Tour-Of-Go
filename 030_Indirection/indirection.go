package main

import "fmt"
import "math"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

// функции, принимающие аргумент определённого типа
// должны получать значение этого типа
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// что касается методов, они могут принимать как
// значение, так и указатель
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// Go интерпретирует v.Scale(2) как (&v).Scale(2)
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p = &Vertex{4, 3}
	// вызов метода p.Abs() интерпретируется как (*p).Abs()
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
