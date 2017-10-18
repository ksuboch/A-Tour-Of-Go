package main

import "fmt"

func main() {
	// отложенное выполнение (до выхода из функции)
	defer fmt.Println("world")

	fmt.Println("Hello")

	fmt.Println("counting...")

	// отложенные вызовы накапливаются в стеке
	// после выхода из функции, вызовы выполняются
	// в порядке LIFO
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
