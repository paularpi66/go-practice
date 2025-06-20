package main

import "fmt"

type Student struct {
	Name string
	Marks int
}

func main () {
	records  := make (map[int]Student)

	records[101] = Student{Name: "Arpita", Marks: 95}
	records[102] = Student{Name: "Arpita", Marks: 75}
	records[103] = Student{Name: "Arpita", Marks: 85}

	for id, student := range records {
		fmt.Println("ID: %d, Name: %s, Marks: %d\n", id, student.Name, student.Marks)
	}

}