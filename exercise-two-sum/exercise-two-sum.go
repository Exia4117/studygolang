package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[target-v] = i
	}
	for j, v := range nums {
		if index, ok := m[v]; ok && m[v] != j {
			return []int{j, index}
		}
	}
	return []int{}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	answer := twoSum(nums, target)
	fmt.Println(answer)
}
