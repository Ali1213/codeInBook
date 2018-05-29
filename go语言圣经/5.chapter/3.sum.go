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

func Max(vals ...int) int {
	fmt.Println(vals)
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	fmt.Println(Max())
}
