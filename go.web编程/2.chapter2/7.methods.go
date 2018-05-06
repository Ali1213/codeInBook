package main

import(
	f "fmt"
	"math"
)

type Rectange struct {
	width,
	height float64
}

type Circle struct {
	redius float64
}

func (r Rectange) area() float64{
	return r.width * r.height
}

func (c Circle) area() float64{
	return c.redius * c.redius * math.Pi
}

func main(){
	r1 := Rectange{12,2}
	r2 := Rectange{9,4}
	c1 := Circle{10}
	c2 := Circle{25}

	f.Println("Area of r1 is: ", r1.area())
	f.Println("Area of r2 is: ", r2.area())
	f.Println("Area of c1 is: ", c1.area())
	f.Println("Area of c2 is: ", c2.area())
}