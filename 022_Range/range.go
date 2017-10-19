package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// возвращается два значения - индекс и элемент
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow1 := make([]int, 10)

	// если нужен только индекс
	// опускается второй аргумент
	for i := range pow {
		pow1[i] = 1 << uint(i) // == 2**i
	}

	// если нужны только значения
	// первый аргумент заменяется на _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
