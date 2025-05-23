package main

import (
	"fmt"
)

func updataName(x *string) {
	*x = "wedge"

}
func main() {
	name := "tifa"
	//updataName(name)
	//fmt.Println("memory address of name is: ", &name)

	m := &name
	fmt.Println("memory address: ", m)
	fmt.Println("value at memory address: ", *m)
	updataName(m)

	fmt.Println(name)

}
