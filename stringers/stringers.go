package main

import "fmt"

//fmt.Println或打印一个变量的值的时候，会判断这个变量是否实现了Stringer接口
//如果实现了，则调用这个变量的String()方法，并将返回值打印到屏幕上

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
	//Sprintf
	//1. 用传入的格式化规则符将传入的变量格式化，(终端中不会有显示)
	//2. 返回为 格式化后的字符串。
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
