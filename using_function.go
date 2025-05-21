package main

import (
	"fmt"
)

func sayGretting(n string) {
	fmt.Println("good morning", n)
}
func sayBye(n string) {
	fmt.Println("good bye", n)
}
func cyclenames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	//sayGretting("mario")
	//sayGretting("Ryan")
	//sayBye("mario")
	cyclenames([]string{"cloud", "fifa", "tifa", "mozilla"}, sayGretting)
	cyclenames([]string{"cloud", "fifa", "tifa", "mozilla"}, sayBye)

}
