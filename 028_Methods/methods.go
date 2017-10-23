package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs - функция со специальным аргументом - приёмником
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

// нельзя объявить метод, тип приемника которого определён
// в другом пакете (это относится и к встроенным типам)
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
