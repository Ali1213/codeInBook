package main

import (
	"fmt"

	"./1.geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
}
