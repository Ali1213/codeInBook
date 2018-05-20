package main

import (
	"fmt"
	"strings"
)

func main() {

	path1 := "/user/logs/aaa"
	path2 := "/user/logs/aaa.csv"
	path3 := "aaa"
	fmt.Println(basename(path1))
	fmt.Println(basename(path2))
	fmt.Println(basename(path3))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > -1 {
		s = s[:dot]
	}
	return s
}
