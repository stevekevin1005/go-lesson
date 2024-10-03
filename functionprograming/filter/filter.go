package main

import (
	"log"
)

func apply(arr []int, mapper func(int) int) []int {
	out := make([]int, len(arr))

	for i, e := range arr {
		out[i] = mapper(e)
	}

	return out
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	out := apply(arr, func(n int) int { return n * n })

	log.Println(out)
}
