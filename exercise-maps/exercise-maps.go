package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	c := strings.Fields(s)
	//fmt.Printf("%q", c)
	for _, w := range c {
		m[w]++
	}
	return m
}

func main() {
	//str := "I ate a donut. Then I ate another donut."
	// fmt.Println(WordCount(str))
	wc.Test(WordCount)
}
