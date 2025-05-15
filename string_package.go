package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "hello there friends"

	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	//fmt.Println(strings.Toupper(greeting))
	//fmt.Println(strings.Index(greeting, "hi"))
	//fmt.Println(strings.Split(greeting, "hi"))

	//the original value is unchanged
	fmt.Println("original string value =", greeting)

}
