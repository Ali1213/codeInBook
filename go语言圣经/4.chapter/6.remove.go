package main

import (
	"fmt"
)

func main() {
	c := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(remove(c, 5))
}

func remove(a []int, pos int) []int {
	copy(a[pos:], a[pos+1:len(a)])
	return a[:len(a)-1]
}
