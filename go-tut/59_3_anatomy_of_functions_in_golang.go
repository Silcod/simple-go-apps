package main

import (
  "fmt"
)
// Multiple return values: In case of multiple returns values but you are only interested in one single value returned by the
// function, you can assign other value(s) to _ (blank identifier)

func addMult(a, b int) (int, int) {
  return a + b, a * b
}

func main() {
  addRes, _ := addMult(2, 5)
  fmt.Println(addRes)
}
