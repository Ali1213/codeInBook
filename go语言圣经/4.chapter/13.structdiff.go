package main

import (
	"fmt"
)

type test struct {
	a, b int
}

func main() {
	a := test{1, 2}
	b := test{1, 2}
	fmt.Println(a == b)
}
