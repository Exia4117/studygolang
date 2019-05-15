package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	//a len = 5 cap = 5 [0 0 0 0 0]
	b := make([]int, 0, 5)
	printSlice("b", b)
	//
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}
