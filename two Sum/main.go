package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))

}

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, v := range numbers {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return []int{}
}
