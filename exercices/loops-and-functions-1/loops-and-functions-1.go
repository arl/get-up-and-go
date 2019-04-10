package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for n := 0; n < 10; n++ {
		z -= (z*z - x) / (z * 2)
	}
	return z
}

func main() {
	var x float64 = 2
	fmt.Println("math.Sqrt: ", math.Sqrt(x))
	fmt.Println("my   Sqrt: ", Sqrt(x))
}
