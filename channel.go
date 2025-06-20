package main

import(
	"fmt"
)

func main() {
	ch := make(chan string)
	//sender go routine
	go func() {
		ch <- "hello go routine"
	}()

	//main goroutine reciver

	msg := <-ch
	fmt.Println(msg)
}