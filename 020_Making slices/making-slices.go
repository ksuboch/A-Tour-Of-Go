package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]int, 5) // len(a)=5
	printSlice("a", a)

	// третий аргумент - вместимость среза
	b := make([]int, 0, 5)
	printSlice("b", b) // len(b)=0, cap(b)=5

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// срезы могут содержать данные любого типа
	// в том числе и другие срезы
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
