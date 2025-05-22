package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"tpffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "_", v)
	}

	//ints as key type
	phonebook := map[int]string{
		234555: "mario",
		234444: "jhon",
		432244: "smith",
		908884: "hurry",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[234555])

	phonebook[432244] = "Browser"
	fmt.Println(phonebook)

	phonebook[908884] = "yoshi"

}
