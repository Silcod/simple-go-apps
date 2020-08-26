/*
  Like JavaScript, functions in Go are first-class citizens. They can be assigned to variables, passed
  as an argument, immediately invoked or deferred for last execution.
*/
package main

import (
  "fmt"
)

func add(a, b int) int64 {
  return int64(a + b) // Note the int64() must be present unless it will give an error
}

func main() {
  result := add(2,3)
  fmt.Println(result)
}
