package main

import (
	"fmt"
)

func main() {
	a := []string{"hello", "", "world"}
	fmt.Println(noempty(a))
}

func noempty(a []string) []string {
	var i int

	for k, v := range a {
		if v != "" {
			a[i] = a[k]
			i++
		}
	}
	return a[:i]
}
