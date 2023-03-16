// The area of Triangle
package main

import "fmt"

func main() {

	var base, height float32

	fmt.Println("Enter the base of triangle: ")
	fmt.Scanf("%f", &base)
	fmt.Println("Enter the height of triangle: ")
	fmt.Scanf("%f", &height)

	area := float32(0.5) * base * height
	fmt.Println("The area of Triangle: ", area)

}
