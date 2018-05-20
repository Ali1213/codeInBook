package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "bac"
	b := "abc"
	c := "acc"
	fmt.Println(isMath(a, b))
	fmt.Println(isMath(a, c))
}

func isMath(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	t1, t2 := strings.Split(s1, ""), strings.Split(s2, "")
	for _, v := range t1 {
		has := false
		for k, v2 := range t2 {
			if v == v2 {
				t2[k] = ""
				has = true
				break
			}
		}
		if !has {
			return false
		}
	}
	for _, v := range t2 {
		if v != "" {
			return false
		}
	}
	return true
}
