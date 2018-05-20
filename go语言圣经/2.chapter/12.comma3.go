package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	a := "123123"
	b := "1212121323"
	c := "121"
	d := "12"
	e := "121212.1323"
	f := "+121212.1323"
	g := "+121212.1323"
	h := "+21212.1323"
	fmt.Println(comma(a))
	fmt.Println(comma(b))
	fmt.Println(comma(c))
	fmt.Println(comma(d))
	fmt.Println(comma(e))
	fmt.Println(comma(f))
	fmt.Println(comma(g))
	fmt.Println(comma(h))
}

func comma(s string) string {
	var b bytes.Buffer
	len := len(s)
	if p := strings.Index(s, "."); p != -1 {
		len = p
	}
	pos := len % 3
	if pos == 0 {
		pos = 3
	}
	if string(s[0]) == "+" || string(s[0]) == "-" {
		if len%3 != 0 && pos == 1 {
			pos += 3
		}
	}
	for k, v := range s {
		if k == pos && k < len {
			b.WriteString(",")
			pos += 3
		}
		b.WriteRune(v)
	}
	return b.String()
}
