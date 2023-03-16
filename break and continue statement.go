package main

import "fmt"
func main(){

	// break, continue statement

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v\n",i)
	}
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			break
		}
		fmt.Printf("%v\n",i)
	}

}