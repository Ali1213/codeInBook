package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := sha256.Sum256([]byte("c"))
	fmt.Println(len(a))
}
