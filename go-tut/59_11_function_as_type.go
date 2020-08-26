package main

import (
  "fmt"
)

// We can create a type alias that will make things simpler. We can rewrite the above program as

func add(a int, b int) int {
  return a + b
}

func subtract(a int, b int) int {
  return a - b
}

type CalcFunc func(int, int) int

func calc(a int, b int, f CalcFunc) int {
  r := f(a, b) // calling add(a,b) or subtract(a, b)
  return r
}

func main() {
  addResult := calc(5, 3, add)
  subResult := calc(5, 3, subtract)
  fmt.Println("5+3 =", addResult)
  fmt.Println("5-3 =", subResult)
}

/*
    Since we understood that a function has its own type, we can declare a variable of type function
    and assign it layer, like following: var add func(int, int) int.
   The syntax above will declare a variable of type function which takes two int arguments and returns
   a int value.
*/
