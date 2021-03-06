package main

import f "fmt"


type person struct {
	name string
	age int
}


func older(p1,p2 person) (person, int) {
	if(p1.age>p2.age){
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}

func main(){
	var tom person
	tom.name,tom.age = "Tom", 19
	bob := person{age:25,name:"Bob"}
	paul := person{"Paul",43}

	tb_older, tb_diff := older(tom,bob)
	tp_older,tp_diff := older(tom,paul)
	bp_older,bp_diff := older(bob,paul)

	f.Printf("Of %s and %s, %s is older by %d years\n",tom.name,bob.name,tb_older.name,tb_diff)
	f.Printf("Of %s and %s, %s is older by %d years\n",tom.name,paul.name,tp_older.name,tp_diff)
	f.Printf("Of %s and %s, %s is older by %d years\n",bob.name,paul.name,bp_older.name,bp_diff)

}