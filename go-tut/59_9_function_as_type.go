package main

import (
  "fmt"
)
/*
  Function as type: a function in Go is also a type. If two function accepts the same parameters and returns the same values,
   then these two functions are of the same type.
*/

func add(a int, b int) int {
  return a + b
}

func subtract(a int, b int) int {
  return a - b
}

func main() {
  fmt.Printf("Type of function add is              %T\n", add)
  fmt.Printf("Type of function subtract is         %T\n", subtract)
}

//So you can see that both add and subtract function has the same type func(int, int) int.
