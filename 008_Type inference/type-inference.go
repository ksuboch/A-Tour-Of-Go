package main

import "fmt"

func main() {
	i := 42           // int
	j := 3.124        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("v is of typre %T\n", i)
	fmt.Printf("v is of typre %T\n", j)
	fmt.Printf("v is of typre %T\n", g)
}
