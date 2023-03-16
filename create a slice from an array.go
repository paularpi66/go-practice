//create a slice from an array:

package main
import ("fmt")

func main() {
  arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice1 := arr1[1:4]
  myslice2 := arr1[2:4]
  myslice3 := arr1[3:5]
  
  fmt.Printf("myslice1: \n",)
  fmt.Printf("myslice = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n\n", cap(myslice1))
  
  fmt.Printf("myslice2: \n",)
  fmt.Printf("myslice = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n\n", cap(myslice2))

  fmt.Printf("myslice3: \n",)
  fmt.Printf("myslice = %v\n", myslice3)
  fmt.Printf("length = %d\n", len(myslice3))
  fmt.Printf("capacity = %d\n\n", cap(myslice3))
}