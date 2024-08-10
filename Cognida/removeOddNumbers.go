package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Original List:", numbers)

	evenNumbers := make([]int, 0, len(numbers))
	for _, n := range numbers {
		if n%2 == 0 {
			evenNumbers = append(evenNumbers, n)
		}
	}

	fmt.Println("List after removing odd numbers:", evenNumbers)
}
