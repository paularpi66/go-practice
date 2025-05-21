package main

import "fmt"

func main() {
	//	x := 0
	//	for x < 5 {
	//		fmt.Println("The value of x :", x)
	//		x++
	//	}
	//

	//names := []string{"John", "Paul", "George", "Ringo"}
	//for i := 0; i < len(names); i++ {
	//	fmt.Println(names[i])
	//}

	names := []string{"John", "Paul", "George", "Ringo"}
	for index, value := range names {
		fmt.Println(index, value)
	}
}
