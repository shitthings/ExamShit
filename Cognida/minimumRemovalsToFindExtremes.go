package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	min := math.MaxInt
	max := math.MinInt
	mini := -1
	maxi := -1

	for i, val := range arr {
		if val < min {
			min = val
			mini = i
		}
		if val > max {
			max = val
			maxi = i
		}
	}

	fsmin := mini + 1
	fsmax := maxi + 1
	femin := n - mini
	femax := n - maxi
	o1 := max(fsmin, fsmax)
	o2 := max(fsmax, fsmin)
	o3 := fsmin + femax
	o4 := fsmax + femin
	result := min(min(o1, o2), min(o3, o4))

	fmt.Println(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
