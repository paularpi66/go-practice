package main

import "fmt"
func main(){

	// series -> 2 + 4 + 6 + 8 + ... N
	var startNumber, endNumber int

	fmt.Printf("Enter the starting number of series: ")
	fmt.Scan(&startNumber)

	fmt.Printf("Enter the last number of series: ")
	fmt.Scan(&endNumber)

	sum := 0
	for i := startNumber; i <= endNumber; i=i+2 {
	  sum = sum + i
	}
	fmt.Printf("sum = %v\n",sum)

}