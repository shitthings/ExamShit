package main

import (
	"fmt"
	"math"
)

func main() {
	var d string
	fmt.Scan(&d)

	n, s, e, w := 0, 0, 0, 0
	for _, dir := range d {
		switch dir {
		case 'N':
			n++
		case 'S':
			s++
		case 'E':
			e++
		case 'W':
			w++
		}
	}

	v := n - s
	h := e - w
	minChanges := int(math.Abs(float64(v))) + int(math.Abs(float64(h)))

	if minChanges%2 == 0 {
		fmt.Println(minChanges / 2)
	} else {
		fmt.Println(-1)
	}
}
