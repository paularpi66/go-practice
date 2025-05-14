package main

import "fmt"

func main() {

	//var ages [3]int = [3]int{20,25,30,35}
	var ages = [3]int{20, 25, 30}
	names := [4]string{"yoshi", "mario", "bowser"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices (use arrays under the hood
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	//slice range
	rangeOne := names[1:3]
	rangeTow := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTow, rangeThree)
	rangeOne = append(rangeOne, "james")
	fmt.Println(rangeOne)

}
