package main

import (
	"fmt"
	"strconv"
)
type Elment interface{}

type List []Elment

type Person struct {
	name string
	age int
}

func (p Person)String()string{
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main(){
	list := make(List,3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"paul",23}
	for _,e := range(list) {
		if value,ok := e.(int); ok {
			fmt.Printf("The type of this is int: %d \n", value)
		}else if value,ok := e.(string); ok {
			fmt.Printf("The type of this is string: %s \n",value)
		}else if value,ok := e.(Person);ok {
			fmt.Printf("The type of this is Person: %s",value)
		}
	}
}

