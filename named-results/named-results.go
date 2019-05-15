package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //返回x和y，均为整型
}

func main() {
	fmt.Println(split(17))
	//7 10
}
