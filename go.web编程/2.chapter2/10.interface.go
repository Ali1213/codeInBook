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
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}


func (h Human) sayHi(){
	fmt.Printf("name: %s, age: %d say hi to you\n",h.name,h.age)
}

func (h Human) sing(lyrics string){
	fmt.Println("La la la la ……", lyrics)
}

// 这里都不能是引用类型
func (e Employee) sayHi(){
	fmt.Printf("Hello, I'm a employee,my name is %s", e.name)
}

type Men interface{
	sayHi()
	sing(lyrics string)
}


func main(){
	mark := Student{Human{"Mark",23,"13701211232"},"MIT",0.00}
	paul := Student{Human{"Paul",26,"4131155445"},"NEW WORK",100}
	sam := Employee{Human{"Sam",33,"121341232"},"Golang",1000}
	Tom := Employee{Human{"Tom",24,"123145513"},"ThinkPad",5000}

	var i Men

	i = mark
	fmt.Println("This is mark,a student")
	i.sayHi()
	i.sing("November rain")

	i=Tom
	fmt.Println("This is tom, a employee")
	i.sayHi()
	i.sing("Born to be wild")

	fmt.Println("Let's user a slice of Men and see what happen")

	x := make([]Men,3)
	x[0],x[1],x[2] = paul,sam,mark
	for _,value := range x {
		value.sayHi()
	}

}