package main

import "fmt"

func add(x, y float32) float32{
  return x + y
}
func sub(x, y float32) float32{
  return x - y
}
func mul(x, y float32) float32{
  return x * y
}
func div(x, y float32) float32{
  return x / y
}

func main(){
  var num1, num2, result float32
  var option string

  i:=true
  for  i==true {
	fmt.Printf("num1 = ")
	fmt.Scan(&num1)

	fmt.Printf("num2 = ")
	fmt.Scan(&num2)

	fmt.Printf("choose an option ( + - * / ) : ")
	fmt.Scan(&option)

	switch option {
	  case "+":
		result = add(num1, num2)
	  case "-":
		result = sub(num1, num2)
	  case "*":
		result = mul(num1, num2)
	  case "/":
		result = div(num1, num2)
	  default:
		fmt.Println("Invalid option")
		continue;
	}
	fmt.Printf("Result = %v\n",result)
  }
}