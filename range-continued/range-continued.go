package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) //uint无符号整型
	}
	for _, value := range pow { //可以通过赋值给 _ 来忽略序号和值
		fmt.Printf("%d\n", value)
	}
	/*
		1
		2
		4
		8
		16
		32
		64
		128
		256
		512
	*/
}
