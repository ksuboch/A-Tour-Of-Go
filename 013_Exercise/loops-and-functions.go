package main

import (
	"fmt"
	"math"
)

// Sqrt - Neuthon sqrt method
func Sqrt(x float64) float64 {
	var zN, zN1 float64 = 1.0, 0.0

	zN1 = zN - ((zN*zN)-x)/(2*zN)

	for math.Abs(zN-zN1) > 0.000001 {
		zN = zN1
		zN1 = zN - ((zN*zN)-x)/(2*zN)
	}
	return zN1
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
