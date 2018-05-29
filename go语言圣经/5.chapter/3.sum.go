package main

import (
	"fmt"
)

func Sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func Max(vals ...int) int, bool {
	if len(vals) < 0 {
		return nil, false
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max, true
}

func main() {
	fmt.Println(Max())
}
