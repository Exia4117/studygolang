package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return x
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
	// fmt.Println(v)
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	//27 >= 20
	// 3 20
}
