package main

import "fmt"

func main() {
	a := []int{34, 5, 5, 32, 34}
	b := []int{4, 3, 3, 5, 0}

	sum := 0
	aSize := len(a)

	for _, index := range b {
		if index < aSize {
			sum += a[index]
		}
	}

	fmt.Println(sum)
}
