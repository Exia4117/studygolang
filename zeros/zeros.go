package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q \n", i, f, b, s)
	//0 0 false ""
	//未赋值的int、float64、bool、string的_零值_
}
