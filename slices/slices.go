package main

import "fmt"

func main() {
	p := []int{2, 3, 7, 11, 13}
	fmt.Println("p ==", p)
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d \n", i, p[i])
	}
	//p == [2 3 7 11 13]
	// p[0] == 2
	// p[1] == 3
	// p[2] == 7
	// p[3] == 11
	// p[4] == 13
}
