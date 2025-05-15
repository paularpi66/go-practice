package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{45, 20, 35, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 35)
	fmt.Println(index)

	names := []string{"yoshi", "maria", "bang", "dong", "ding"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "maria"))
}
