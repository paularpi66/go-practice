package main

import (
	"fmt"
	"math"
)

//interface
type Shape interface{
	Area() float64
}

//Struct 1: circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Struct 2: Rectangle
type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r. Height
}

//Function that takes a Shape interface
func printArea( s Shape) {
	fmt.Println("Area is:", s.Area())
}
funce main(){
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 3}

	printArea(c)
	printArea(r)
}