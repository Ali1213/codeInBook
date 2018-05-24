package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	utfLens := [utf8.MaxRune+1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r,n,err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if e != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n",err)
			os.Exit(1)
		}

		if  r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utfLens[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nLens\tcount\n")
	for i, n := range utfLens {
		fmt.Printf("%d\t%d\n", i, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}
