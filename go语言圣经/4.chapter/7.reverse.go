package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reverse(&a))
}

func reverse(a *[]int) *[]int {
	mid := len(*a) / 2
	for i, j := 0, len(*a)-1; i <= mid; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
	return a
}
