package main
import "fmt"

func ChangeValue( value int){
	value = 30
}
func CallByReference( ptr *int){
	*ptr = 20
}

func main(){
	x := 10
	fmt.Println(x)

	ChangeValue(x)
	fmt.Println(x)

	CallByReference(&x)
	fmt.Println(x)
}