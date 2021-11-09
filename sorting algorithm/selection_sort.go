package main

import "fmt"

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
