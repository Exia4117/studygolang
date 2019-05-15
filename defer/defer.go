package main

import "fmt"

func main() {
	defer fmt.Println("world") //延迟函数的执行直到上层函数返回
	fmt.Println("hello")
}
