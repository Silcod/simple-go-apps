package main

import (
  "fmt"
)

func main() {
  sum := func(a int, b int) int {
    return a + b
  }(3, 5)

  fmt.Println("5 + 3 =", sum)
}

/*
  In the above program, look at the function definition. The first part from func to } defines the
  function while later (3, 5) executes it. Hence sum is the value returned by a function execution.
  Hence the above program yields the following result
*/
