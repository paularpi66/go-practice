package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	name, _ := r.ReadString('\n')

	return strings.TrimSpace(name), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	//fmt.Print("Create a new Bill: ")
	//name, _ := reader.ReadString('\n')
	//name = strings.Trimepace(name)

	name, _ := getInput("Bill name: ", reader)

	b := newBill(name)
	fmt.Println("create a new Bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option: (a - add iteams, s - save bill, t - add tip ", reader)
	fmt.Println(opt)
}
func main() {
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}
