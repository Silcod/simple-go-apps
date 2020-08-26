package main

import (
  "fmt"
)
// Named return values: Named return values are a great way to explicitly mention return variables in the function definition itself.

func addMult(a, b int) (add int, mul int) { // you can also use it as (add, mul int)
  add = a + b
  mul = a * b

  return // This is necessary
}

func main() {
  addRes, multRes := addMult(2, 5)
  fmt.Println(addRes, multRes)
}
