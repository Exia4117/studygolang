package main

import "fmt"

func main() {
	v := 0.867 + 0.5i
	fmt.Printf("v is of type %T \n", v)
	//未指明类型数字常量，类型取决于常量精度
	//%T为类型占位符
}
