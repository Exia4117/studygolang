package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	//Hello World
	fmt.Println(a)
	//[Hello World]
	//数组的len和cap是永远相等的，并且是在定义的时候就已经指定的，不能改变
}
