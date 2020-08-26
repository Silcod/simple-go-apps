package main

import (
  "fmt"
)
// Immediately-invoked function expression (IIFE):  In Go, we can create an anonymous function that can be
// defined and executed at the same time.
func main() {
  var add = func(a int, b int) int {
    return a + b
  }

  fmt.Println("5 + 3 =", add(5,3))
}

/*
  Where, add is an anonymous function. Some can argue that it’s not truly anonymous because we can still
  refer to the add function from anywhere in main function (in other cases, from anywhere in the program).
  But not in the case when a function is immediately invoked or executed. Let’s modify the previous example.
*/
