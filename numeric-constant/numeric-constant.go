package main

import "fmt"

const (
	Big   = 1 << 100  //位运算，左移100位，即乘2的100次方(2的100次方)
	Small = Big >> 99 //位运算，右移99位，即除以2的99次方(Small = 2)
) //常量不能用:=定义

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	//21
	fmt.Println(needFloat(Small))
	//0.2
	fmt.Println(needFloat(Big))
	//1.2676506002282295e+29
	// fmt.Println(needInt(Big))
}
