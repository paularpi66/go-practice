package main

import "fmt"

func multipler(factor int) func(int) int {
	return func(x int) int {
		return X * factor
	}
}

func main () {
	times2 := multipler (2)
	times3 := multiplier(3)

	fmt.Println(times2(5)) //10(5 X 2)
	fmt.Println(times3(4)) //12 (4 X 3)
}
