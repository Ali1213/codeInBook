package main 

import (
	"fmt"
	"strings"
	"os"
)

func main(){
	s := strings.Join(os.Args[1:],"   ")
	fmt.Println(s)
	
}