package main

import(
  "fmt"
)

func main() {
  var a = 5
  var p = &a
  fmt.Printf("Address of var a: %p\n", p)
  fmt.Printf("Value of var a: %v\n", *p)

  myValue := "Hi"
  var myVabu = &myValue
  fmt.Println(myVabu)
  fmt.Println(*myVabu)

  myLove := "Beauty"
  fmt.Println(&myLove)
  fmt.Println(*&myLove)

  /*
    Declaring a pointer: Below is a pointer of type string which can store only the memory addresses of string variables:
      var s *string
      var p *int
  */
}
