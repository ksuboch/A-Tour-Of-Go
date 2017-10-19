package main

import "fmt"

// функция возвращает числа Фибоначчи
func fibonacci() func() int {

	var fib1, fib2 = 0, 0

	return func() int {
		switch fib2 {
		case 0:
			fib1, fib2 = 0, 1
		default:
			fib1, fib2 = fib2, fib1+fib2
		}
		return fib1
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
