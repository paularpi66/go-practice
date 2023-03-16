// If else odd/even problem
package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter any integer number: ")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd\n")
	}
}
