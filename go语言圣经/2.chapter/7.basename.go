package main

import . "fmt"

func main() {
	path1 := "/user/logs/aaa"
	path2 := "/user/logs/aaa.csv"
	path3 := "aaa"
	Println(basename(path1))
	Println(basename(path2))
	Println(basename(path3))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
