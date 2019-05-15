package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z := float64(1)
	// fmt.Println(math.Abs(z - math.Sqrt(x)))
	for math.Abs(z-math.Sqrt(x)) > 0.0000001 {
		z = z - (z*z-x)/(2*z)
		// return z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt2(2))
	fmt.Println(math.Sqrt(2))
}
