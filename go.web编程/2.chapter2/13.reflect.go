package main

import (
	"fmt"
	"reflect"
)

func main(){
	i := 3.4

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	fmt.Println(t)
	fmt.Println(v.Type())

}