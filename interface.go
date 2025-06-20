package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct{}
func (p Person) Speak() string {
	return "Hello, I'm a person!"
}

type Dog srruct{}
func (d Dog) Speak() string {
	return "Woof!"
}

func saySomething (s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	p := Person{}
	d := Dog{}


	saySomething(p)
	saySomething(d)
}