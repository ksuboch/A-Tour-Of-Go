package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// T реализует интерфейс I
// без специального указания этого
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
