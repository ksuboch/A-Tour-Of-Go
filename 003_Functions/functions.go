package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func pAdd(x, y *int) int {
	return *x + *y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	a, b := new(int), new(int)
	*a, *b = 2, 3
	fmt.Println(pAdd(a, b))
	fmt.Println(swap("world", "Hello"))
}
