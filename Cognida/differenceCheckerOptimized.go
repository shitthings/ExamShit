package main

import (
	"fmt"
)

func checkDifferences(arr []int, k int) []int {
	n := len(arr)
	result := make([]int, 0, n)
	set := make(map[int]struct{}, n)

	for _, num := range arr {
		set[num] = struct{}{}
	}

	for _, current := range arr {
		if _, exists := set[current+k]; exists || _, exists := set[current-k]; exists {
			result = append(result, 1)
		} else {
			result = append(result, 0)
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
