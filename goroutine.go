package main

import (
    "fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go sayHello() //goroutine
	for i := 0; i < 3; i++ {
		fmt.Println("Main")
		time.Sleep(500 * time.Millisecond)
	}
}