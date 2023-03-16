package main

import "fmt"

func main() {

	var celsius, temperature float32
	fmt.Print(`Please Enter the temerature in fahrenheit:  
	`)
	fmt.Scanf("%f", &temperature)
	celsius = (temperature - 32) * 5 / 9
	fmt.Println("Temperature in celcious: ", celsius)
}
