package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	s = appendInt(s, 4)
	fmt.Println(s)
}

func appendInt(x []int, y int) []int {
	var a []int
	aLen := len(x) + 1
	if aLen < cap(x) {
		a = x[:aLen]
	} else {
		aCap := aLen
		if aCap < 2*len(x) {
			aCap = 2 * len(x)
		}
		a = make([]int, aLen, aCap)
		copy(a, x)
	}
	a[len(x)] = y
	return a
}
