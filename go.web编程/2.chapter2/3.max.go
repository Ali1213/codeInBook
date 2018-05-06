package main

import "fmt"

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func main(){
	x := 2
	y := 3
	z := 4

	max_xy := max(x,y)
	max_xz := max(x,z)

	fmt.Printf("max (%d, %d) = %d\n", x,y,max_xy)
	fmt.Printf("max (%d, %d) = %d\n", x,z,max_xz)
	fmt.Printf("max (%d, %d) = %d\n", y,z,max(y,z))
}