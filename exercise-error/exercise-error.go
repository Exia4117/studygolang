package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	z := float64(1)
	if x >= 0 {
		for i := 0; i < 10; i++ {
			z = z - (z*z-x)/(2*z)
		}
	} else {
		return 0.0, ErrNegativeSqrt(x)
	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can't sqrt negative number : %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
