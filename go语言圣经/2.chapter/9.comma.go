package main

import (
	"fmt"
)

func main() {
	p := "123354535"
	fmt.Println(comma(p))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
