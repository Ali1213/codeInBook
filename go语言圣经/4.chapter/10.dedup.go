package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	store := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := input.Text()
		if !store[t] {
			store[t] = true
			fmt.Println(t)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
