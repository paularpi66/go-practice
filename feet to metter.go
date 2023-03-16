package main

import "fmt"

func main() {
	var n int
	fmt.Print("Please Enter the height in Centemeters:  ")
	fmt.Scanf("%d", &n)
	inches := 0.394 * float32(n)
	feet := 0.328 * float32(n)
	meter := 0.3048 * feet
	fmt.Println("The Length in Inches:", inches)
	fmt.Println("The length in feet: ", feet)
	fmt.Println("The length in meter: ", meter)
}
