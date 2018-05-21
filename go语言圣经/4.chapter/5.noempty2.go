package main

import (
	"fmt"
)

func main() {
	a := []string{"hello", "", "world"}
	fmt.Println(noempty(a))
}

func noempty(a []string) (s []string) {
	s = a[:0]
	for _, v := range a {
		if v != "" {
			s = append(s, v)
		}
	}
	return
}
