package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "12"
	num2 := "15"

	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	result := n1 * n2

	resultString := strconv.Itoa(result)
	fmt.Println("Multiplication result:", resultString)
}
