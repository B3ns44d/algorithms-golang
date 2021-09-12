package main

import (
	"fmt"
)

func main() {
	var a []int
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println(a)
	reverse(a)
	fmt.Println(a)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
