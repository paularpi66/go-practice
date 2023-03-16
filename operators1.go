package main

import "fmt"

func main() {
	var x = 5
	var y = 10

	x = y
	fmt.Println(x)

	x += 1
	fmt.Println(x)

	x -= 1
	fmt.Println(x)

	x *= 1
	fmt.Println(x)

	x /= 1
	fmt.Println(x)

	x += y
	fmt.Println(x)

	x -= y
	fmt.Println(x)

	x *= y
	fmt.Println(x)

	x /= y
	fmt.Println(x)
}
