package main

import "fmt"

func main() {
	var z []int
	fmt.Println(len(z), cap(z), z)
	//0 0 []
	if z == nil {
		fmt.Println("nil!")
	}
}
