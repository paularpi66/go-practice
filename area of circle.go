// The area of Circle
package main

import "fmt"

func main() {

	var radius float32

	fmt.Println("Enter the radius of circle: ")
	fmt.Scanf("%f", &radius)

	area := 3.1416 * radius * radius
	fmt.Println("The area of Circle: ", area)

}
