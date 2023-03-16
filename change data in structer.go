package main

    import "fmt"

    type Student struct {
      name string
      age int
      id int
    }

    func (x *Student) increaseAge(val int) {
      x.age += val
    }

    func main() {
      s1 := Student{"Anisul Islam", 32}
      s1.increaseAge(1)
      fmt.Println(s1.age)
    }
