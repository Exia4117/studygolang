package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	//a len = 5 cap = 5 [0 0 0 0 0]
	b := make([]int, 0, 5) // 参数2为长度，参数3为容量
	printSlice("b", b)
	//b len = 0 cap = 5 []
	c := b[:2]
	printSlice("c", c)
	//c len = 2 cap = 5 [0 0]
	/*
		容量的改变规则为原容量值减掉起始下标
	*/
	d := c[2:5]
	printSlice("d", d)
	//d len = 3 cap = 3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}
