package main
import "fmt"

func main() {
	var number int
	fmt.Println("Please  eneter a number: ")
	fmt.Scanf("%d", &number)

	switch number {
	case 1, 2, 3, 4, 5:
		fmt.Println("Primary")
	case 6, 7, 8, 9, 10:
		fmt.Println("Secondary")
	case 11, 12:
		fmt.Println("Higher Secondary")
	default: 
	    fmt.Println("Number is invalid")	
	}

}