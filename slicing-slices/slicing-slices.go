package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)
	//p == [2 3 5 7 11 13]
	fmt.Println("p[1:4] == ", p[1:4]) //从_左边下标_到_右边下标-1_，包括两端
	//p[1:4] == [3 5 7]
	fmt.Println("p[1:1] ==", p[1:1]) //为空
	//p[1:1] == []

	//省略下标代表从0开始
	fmt.Println("p[:3] ==", p[:3])
	//p[:3] == [2 3 5]

	//省略上标代表到len(s)结束
	fmt.Println("p[4:] ==", p[4:])
	//p[4:] == [11 13]

}
