package main

import (
	"fmt"
)

func checkDifferences(arr []int, k int) []int {
	n := len(arr)
	result := make([]int, n)
	set := make(map[int]struct{})

	for _, num := range arr {
		set[num] = struct{}{}
	}

	for i, current := range arr {
		if _, exists := set[current+k]; exists || _, exists := set[current-k]; exists {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return result
}

func main() {
	var n, k int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&k)

	result := checkDifferences(arr, k)

	for _, res := range result {
		fmt.Print(res, " ")
	}
}
