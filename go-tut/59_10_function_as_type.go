package main

import (
  "fmt"
)

func add(a int, b int) int {
  return a + b
}

func subtract(a int, b int) int {
  return a - b
}

/*
  In the program below, we have defined a function calc which takes to int arguments a & b and third
   function argument f of type func(int, int) int. Then we are calling f function with a and b as arguments.
*/

func calc(a int, b int, f func(int, int) int) int {
  r := f(a, b)
  return r
}

func main() {
  addResult := calc(5, 3, add)
  subResult := calc(5, 3, subtract)
  fmt.Println("5+3 =", addResult)
  fmt.Println("5-3 =", subResult)
}
