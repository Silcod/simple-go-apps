package main

import (
  "fmt"
)

/* Function as value (anonymous function): a function in Go can also be a value. This means
  you can assign a function to a variable.*/

var add = func(a int, b int) int {
  return a + b
}

func main() {
  fmt.Println("5 + 3 =", add(5, 3))
}
/*
  In the above program, we have created a global variable add and assigned a newly created function to it.
  We have used Go’s type inference to get the type of anonymous function (as we haven’t mentioned the type of add).
   In this case, add is an anonymous function as it was created from a function that doesn’t have a name.
*/
