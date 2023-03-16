//pointer first program

package main
import "fmt"

func main() {
	var p * int
	x := 10
	p = &x
	fmt.Println("the value of x: ", x)
	fmt.Println("The address of the x: ", &x)
	fmt.Println("the value of p: ", p)
	fmt.Println("the value of *p: ", *p)
}