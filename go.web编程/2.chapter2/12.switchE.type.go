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
		switch value := e.(type){
		case int :
			fmt.Printf("The type of this is int: %d \n", value)
		case string:
			fmt.Printf("The type of this is string: %s \n",value)
		case Person:
			fmt.Printf("The type of this is Person: %s \n",value)
		}
	}

}

