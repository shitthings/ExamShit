package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 6, 8, 12, 24, 16}
	s := make(map[int]struct{}, len(n))
	for _, num := range n {
		s[num] = struct{}{}
	}

	cnt := 0
	for _, num := range n {
		if _, exists := s[num*2]; exists {
			cnt++
		}
	}

	fmt.Println(cnt)
}
