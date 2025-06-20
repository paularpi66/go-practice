package main

import "fmt"

func main() {

	products := map[string]float64{
		"sugar": 80.5,
		"rice": 65. 0,
		"Lentill": 120.0,
		"oil": 150.74,
	}
	fmt.Println("Price of Oil: ", products["oil"], "Taka")

	products["flour"] = 55.0

	fmt.Prinln("\nAll products prices: ")
	for item, price := rane products {
		fmt.Println("%s: %2f Taka\n", item, price)
	}

	delete(products, "sugar")

	total := 0.0
	for _, price := range products {
		total += price
	}
	fmt.Prinln("\nTotal price: 52f Taka\n", total)
}