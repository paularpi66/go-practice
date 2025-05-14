package main

import "fmt"

func main() {

	age := "20"
	name := "Arpita"
	//print
	fmt.Print("Hello")
	fmt.Print("BD \n")
	fmt.Print("new line \n")

	//println
	fmt.Println("Hello World")
	fmt.Println("HELLO NINJAS")
	fmt.Println("my age is", age, "and my name is", name)

	//printf (formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n ", age, name)
	fmt.Printf("my age is %q and my name is %q \n ", age, name)
	fmt.Println("age is of type %T", age)
	fmt.Printf("you scored %f points! \n,225.55")
	fmt.Println("you scored %0.1f points! \n,225.55")

	//Sprintf ( save formatted strings)
	var str = fmt.Sprint("my age is %v and my name is %v \n ", age, name)
	fmt.Println("the saved string is: ", str)

}
