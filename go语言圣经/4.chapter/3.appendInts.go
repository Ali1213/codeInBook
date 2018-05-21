package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	fmt.Println(appendInts(a, b))
}

func appendInts(a, b []int) []int {
	var z []int
	tLen := len(a) + len(b)
	if cap(a) > tLen {
		z = a[:len(a)]
	} else {
		zCap := 2 * tLen
		z = make([]int, tLen, zCap)
		copy(z, a)
	}
	copy(z[len(a):], b)
	return z
}
