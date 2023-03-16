package main
import "fmt"

func main(){
var n, i int
fmt.Println("Please enter a number:")
fmt.Scanf("%d", &n)

for i = n; i <= 10; i++{
	if i % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("odd")
	}
 }
}
