package main

import (
	"fmt"
)

func main() {
	x := map[string]int{
		"make": 1,
		"job":  2,
	}
	y := map[string]int{
		"make": 1,
		"job":  2,
	}
	fmt.Println(equal(x, y))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for kx, vx := range x {

		if vy, ok := y[kx]; !ok || vy != vx {
			return false
		}
	}
	return true
}
