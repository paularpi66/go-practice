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

    func displayInfo(s student) {
		fmt.Println("id = %v, name = %v, age = %v, department = %v, grade = %v, address = %v", s.id, s.name, s.age, s.department, s.grade, s.address)
    }
	
	func main() {
	kona := student{66, "kona", 22, "CSE", 3.64, "Dhaka"}
	sharmin := student{65, "sharmin", 24, "CSE", 3.00, "Dhaka"}
	
	fmt.Println("Kona's Details:- ")
    displayInfo(kona)
    
    fmt.Println("Sharmin's Details:- ")
    displayInfo(sharmin)
	
}