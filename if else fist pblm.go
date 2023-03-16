// If else easy problem
package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter number: ")
	fmt.Scanf("%d", &number)

	if number > 0 {
		fmt.Println("Positive number")
	} else if number < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Zero\n")
	}
}
