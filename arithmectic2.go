// Arrithmetic operators
package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Println("Enter your 1st Number: ")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter your 2nd Number: ")
	fmt.Scanf("%d", &num2)

	fmt.Println("The Value of Sum: ", num1+num2)
	fmt.Println("The Value of Sub: ", num1-num2)
	fmt.Println("The Value of mul: ", num1*num2)
	fmt.Println("The Value of div: ", float32(num1)/float32(num2))

}
