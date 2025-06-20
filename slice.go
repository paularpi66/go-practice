package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Prinln(numbers)  //[10 20 30 40 50]
	fmt.Println(numbers[1:4]) //[20 30 40] -> 1 theke 3 porjonto
}