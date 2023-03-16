// If else easy problem
package main

import "fmt"

func main() {
	var number1, number2 int
	fmt.Println("Enter any integer 2 numbers: ")
	fmt.Scanf("%d %d", &number1, &number2)

	if number1 > number2 {
		fmt.Println("Largest Number = ", number1)
	} else if number1 < number2 {
		fmt.Println("Largest Number = ", number2)
	} else {
		fmt.Println("Numbers are equal")
	}
}
