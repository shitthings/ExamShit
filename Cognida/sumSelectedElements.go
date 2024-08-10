package main

import "fmt"

func main() {
	a := []int{34, 5, 5, 32, 34}
	b := []int{4, 3, 0}

	sum := 0
	for _, index := range b {
		sum += a[index]
	}

	fmt.Println(sum)
}
