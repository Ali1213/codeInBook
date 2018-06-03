package main

import (
	"fmt"

	. "./1.bytecounter"
)

var c ByteCounter

func main() {
	c.Write([]byte("hello"))

	fmt.Println(c)
	c = 0

	var name = "Dolly"

	fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c)
}
