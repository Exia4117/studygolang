package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The Value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The Value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"])
	//The Value: 0
	//通过双赋值检测某个键存在：
	v, ok := m["Answer"]
	//如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 v 是 map 的元素类型的零值。
	fmt.Println("The Value:", v, "Present?", ok)
	//The Value: 0 Present? false
}
