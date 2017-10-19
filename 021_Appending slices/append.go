package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append работает с нулевыми срезами
	s = append(s, 0)
	printSlice(s)

	// при необходимости, срез увеличивается
	s = append(s, 1)
	printSlice(s)

	// можно добавить несколько элементов
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
