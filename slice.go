//This example shows how to create slices using the []datatype{values} format:

package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)
  fmt.Println(myslice1 == nil)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
  fmt.Println(myslice2 == nil)

  myslice3 :=[]string{"Hello", "world"}
  fmt.Println(len(myslice3))
  fmt.Println(cap(myslice3))
  fmt.Println(myslice3)
  fmt.Println(myslice3 == nil)

  var myslice4 []string
  fmt.Println(len(myslice4))
  fmt.Println(cap(myslice4))
  fmt.Println(myslice4)
  fmt.Println(myslice4 == nil)
}