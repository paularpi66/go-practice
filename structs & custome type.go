package main

import "fmt"

type bill struct {
	name map[string]float64
	tip  float64
}

// make new bills
func newBill(name string) bill {
b:
	bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}
func main() {

}
