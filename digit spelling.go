  // a go program to spell digits with if, else if, else
  package main

  import "fmt"
  func main(){
	  var digit int
	  fmt.Printf("Enter a digit (0-9): ")
	  fmt.Scanf("%v",&digit)
  
	  if (digit == 0){
			  fmt.Printf("Zero\n")
	  }	else if (digit == 1){
			  fmt.Printf("One\n")
	  }	else if (digit == 2){
			  fmt.Printf("Two\n")
	  }	else if (digit == 3){
			  fmt.Printf("Three\n")
	  }	else if (digit == 4){
			  fmt.Printf("Four\n")
	  }	else if (digit == 5){
			  fmt.Printf("Five\n")
	  }	else if (digit == 6){
			  fmt.Printf("Six\n")
	  }	else if (digit == 7){
			  fmt.Printf("Seven\n")
	  }	else if (digit == 8){
			  fmt.Printf("Eight\n")
	  }	else if (digit == 9){
			  fmt.Printf("Nine\n")
	  }	else {
		  fmt.Printf("Not a digit\n")
	  }
  }