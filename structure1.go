package main
import "fmt"

type student struct{
	id int
	name string
	age int
	department string
	grade float32
	address string	
}

func main() {
kona := student{66, "kona", 22, "CSE", 3.64, "Dhaka"}
sharmin := student{65, "sharmin", 24, "CSE", 3.00, "Dhaka"}

fmt.Println("kona pauls details:- ", kona)
fmt.Println("Sharmin's details:- ", sharmin)



//individual print
fmt.Println("\n \n## Kona's details ##")
fmt.Println("ID: ", kona.id)
fmt.Println("Name: ", kona.name)
fmt.Println("Age: ", kona.age)
fmt.Println("Department: ", kona.department)
fmt.Println("Grade: ", kona.grade)
fmt.Println("Address: ", kona.address)
}