package main
import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
)
func main(){
	counts := make(map[string]int)
	files := os.Args[1:]

	for _,file := range files {
		data, err := ioutil.ReadFile(file)
		if(err != nil){
			fmt.Fprintf(os.Stderr,"dup3: %v\n", err)
			continue
		}
		for _,line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n>1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}