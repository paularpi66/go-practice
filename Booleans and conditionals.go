package main

import (
	"fmt"
)

func main() {
	age := 25
	fmt.Println(age)
	fmt.Println(age <= 40)
	fmt.Println(age >= 40)
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 20)
	fmt.Println(age != 50)

	if age < 10 {
		fmt.Println("age is less then 10")
	} else if age < 20 {
		fmt.Println("age is less then 20")
	} else {
		fmt.Println("age is not less then 25")
	}

	names := []string{"John", "Paul", "George", "Ringo"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		fmt.Println("The vale at pos %v is %v", index, value)
	}
}
