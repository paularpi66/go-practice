package main

import "fmt"

func main() {

	// variable declaration

	var fullName, country string
	var age int
	var gpa float32

	//vriable initialization

	fullName = "Kona Paul"
	country = "Bangladesh"
	age = 22
	gpa = 3.92

	fmt.Println(fullName, "is a Student")
	fmt.Println(fullName, "is a", age, "years old")
	fmt.Println("Her GPA is", gpa, "out of 4")
	fmt.Println(fullName, "is originally from", country)

}
