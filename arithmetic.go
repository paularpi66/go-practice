// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	//	var fullName string
	var num1, num2 int

	//fmt.Println("Enter your Name: ")
	//fmt.Scan(&fullName)

	fmt.Println("Enter your 1st Number: ")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter your 2nd Number: ")
	fmt.Scanf("%d", &num2)
	sum := num1 + num2

	sub := num1 - num2

	mul := num1 * num2

	div := num1 / num2

	fmt.Println("The Value of Sum: ", sum)
	fmt.Println("The Value of Sub: ", sub)
	fmt.Println("The Value of mul: ", mul)
	fmt.Println("The Value of div: ", div)

}
