package main

import "fmt"

func main() {
	var a []int
	printSlice("a", a)
	//a len = 0 , cap = 0 []

	//append works on nil slice
	a = append(a, 0)
	printSlice("a", a)
	//a len = 1 , cap = 1 [0]

	//the slice grows as needed
	a = append(a, 1)
	printSlice("a", a)
	//	a len = 2 , cap = 2 [0 1]

	//we can add more than one element at a time
	a = append(a, 2, 3, 5)
	printSlice("a", a)
	//a len = 5 , cap = 6 [0 1 2 3 4]
	/*
		当元素被添加到切片时,当原始容量是偶数时,切片的容量加倍.但是当原始容量是奇数时,容量增加1然后加倍
	*/
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len = %d , cap = %d %v\n", s, len(x), cap(x), x)
}
