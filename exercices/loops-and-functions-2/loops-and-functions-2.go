package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	lastz := z
	n := 0

	for {
		z -= (z*z - x) / (z * 2)
		if z == lastz {
			break
		}
		lastz = z
		n++
	}
	fmt.Println(n, "cycles")
	return z
}

func main() {
	var x float64 = 30000000.0
	fmt.Println("math.Sqrt: ", math.Sqrt(x))
	fmt.Println("my   Sqrt: ", Sqrt(x))
}
