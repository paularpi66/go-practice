  // a go program to spell digits with switch statements
  package main

  import "fmt"
  func main(){
	  var digit int
	  fmt.Printf("Enter a digit (0-9): ")
	  fmt.Scanf("%v",&digit)
  
	  switch digit {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("one")
		case 2: 
			fmt.Println(" two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5:
			fmt.Println("five")
		case 6:
			fmt.Println("six")
		case 7:
			fmt.Println("seven")
		case 8:
			fmt.Println("eight")
		case 9:
			fmt.Println("nine")
	    default:
			fmt.Println("Invalid digit")
		 
	  }
  }