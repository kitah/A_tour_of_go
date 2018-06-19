package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const EPS = 10e-6
	z := 1.0
	prev := 0.0
	for math.Abs(prev-z) > EPS {
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
