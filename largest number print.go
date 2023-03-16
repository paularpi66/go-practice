package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Enter three Numbers:")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if a == b { //a==b if
		if a == c { //a==c
			fmt.Println("All numbers are Equal")
		} else if a < c { //a<c
			fmt.Println("%v is the Largest Number", c)
		} else { //a>c  = a/b
			fmt.Println("%v and %v is the Largest Number", a, b)
		}
	} else if a < b { //a<b
		if b == c { //b=c
			fmt.Println("%v and %v is the Largest Number is ", b, c)
		} else if b < c { //b<c
			fmt.Println("%v is the Largest Number", c)
		} else { //b>c
			fmt.Println("%v is the largest Number", b)
		}
	} else { //a>b
		if a == c { //a==c
			fmt.Println("%v and %v is the Largest Number is ", a, c)
		} else if a < c { //a<c
			fmt.Println(" %v is the largest Number", c)
		} else { //a>c
			fmt.Printf(" %v is the largest Number", a)
		}

	}

}
