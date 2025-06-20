package main

import (
	"fmt"
	"sync"
	"time"
)

func worker (id int, wg *sync.WaitGroup) {
	defer wg.Done() // signals that this goroutine is done
	fmt.Prinln("Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}


func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)   // increase counter for each goroutine
		go worker (i, &wg)  //launch goroutine
	}
	wg.Wait() //wait for all goroutines to finished
	fmt.Println("all workers finished")
}