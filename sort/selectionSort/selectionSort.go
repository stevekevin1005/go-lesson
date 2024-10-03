package main

import "fmt"

func Selection_Sort(array []int, size int) []int {
	var min_index int
	var temp int
	for i := 0; i < size-1; i++ {
		min_index = i
		// Find index of minimum element
		for j := i + 1; j < size; j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		temp = array[i]
		array[i] = array[min_index]
		array[min_index] = temp
	}
	return array
}
func main() {
	var num = 7
	array := []int{2, 4, 3, 1, 6, 8, 5}
	fmt.Println(Selection_Sort(array, num))
}
