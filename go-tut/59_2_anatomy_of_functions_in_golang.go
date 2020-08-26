package main

import (
  "fmt"
)
// Unlike other programming languages, Go can return multiple values from the function

func addMult(a, b int) (int, int) {
  return a + b, a * b
}

func main() {
  addRes, addMult := addMult(2,5)
  fmt.Println(addRes, addMult)
}
