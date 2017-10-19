package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// срез - указатель на массивы
	// он не хранит никаких данных
	var s []int = primes[1:4]
	fmt.Println(s)

	var k []int = primes[0:5]
	fmt.Println(k)

	names := [4]string{
		"John",
		"Paul",
		"Geogie",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// изменение среза приводит к
	// модификации массива
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// литерал среза такой-же как для массива
	// но без указания размера
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	m := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(m)

	n := []int{2, 3, 5, 7, 11, 13}

	n = n[1:4]
	fmt.Println(n)

	// при создании среза можно опустить верхнюю
	// или нижнюю границу
	n = n[:2]
	fmt.Println(n)

	n = n[1:]
	fmt.Println(n)

	p := []int{2, 3, 5, 7, 11, 13}
	printSlice(p)

	p = p[:0]
	printSlice(p)

	p = p[:4]
	printSlice(p)

	p = p[2:]
	printSlice(p)

	// срез нулевой длины и вместимости
	var t []int
	fmt.Println(t, len(t), cap(t))
	if t == nil {
		fmt.Println("nil!")
	}
}

// len - размер среза (кол-во элементов, которые он содержит)
// cap - вместимость среза (кол-во элементов в низлежащем массиве)
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
