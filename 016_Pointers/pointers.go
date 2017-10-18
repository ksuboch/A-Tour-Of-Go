package main

import "fmt"

func main() {
	i, j := 42, 2701

	// & возвращает указатель на операнд
	p := &i         // point to i
	fmt.Println(*p) // read the i through the pointer

	// * дает доступ к значению по указателю
	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 27   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
