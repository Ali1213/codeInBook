package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) sayHi(){
	fmt.Printf("name: %s, age: %d say hi to you\n",h.name,h.age)
}


func main(){
	mark := Student{Human{"Mark",23,"13701211232"},"MIT"}
	sam := Employee{Human{"Sam",33,"121341232"},"Golang" }
	mark.sayHi()
	sam.sayHi()
}
