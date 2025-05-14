package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{20,25,30,35}
	var ages = [3]int{20, 25, 30}
	names := [4]string{"yoshi", "mario", "bowser"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

}
