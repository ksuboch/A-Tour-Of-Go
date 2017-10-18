package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func pAdd(x, y *int) int {
	return *x + *y
}

func main() {
	fmt.Println(add(42, 13))
	a, b := new(int), new(int)
	*a, *b = 2, 3
	fmt.Println(pAdd(a, b))
}
