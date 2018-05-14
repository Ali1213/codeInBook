package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0

	fmt.Printf("%g℉ or %g ℃\n", freezingF, fToc(freezingF))
	fmt.Printf("%g℉ or %g ℃\n", boilingF, fToc(boilingF))
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
