package main

import "fmt"

func fibonacci() func() int {
	fir, sec := 0, 1
	fmt.Println(fir)
	fmt.Println(sec)
	return func() int {
		sum := fir + sec
		fir = sec
		sec = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
