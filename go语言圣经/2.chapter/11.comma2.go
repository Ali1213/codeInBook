package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := "123123"
	b := "1212121323"
	c := "121"
	d := "12"
	fmt.Println(comma(a))
	fmt.Println(comma(b))
	fmt.Println(comma(c))
	fmt.Println(comma(d))
}

func comma(s string) string {
	var b bytes.Buffer
	len := len(s)
	pos := len % 3
	if pos == 0 {
		pos = 3
	}
	for k, v := range s {
		if k == pos {
			b.WriteString(",")
			pos += 3
		}
		b.WriteRune(v)
	}
	return b.String()
}
