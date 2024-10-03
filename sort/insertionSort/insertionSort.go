package main

import "fmt"

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func main() {
	fmt.Println(insertionSort([]int{5, 3, 2, 1, 0, 4}))
}
