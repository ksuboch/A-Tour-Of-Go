package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// в языке Go можно создать метод
// который может быть вызван с
// нулевым указателем в качестве
// аргумента
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}
