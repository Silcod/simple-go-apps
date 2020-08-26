package main

import (
  "fmt"
)
/* Recursive Function: A function is called recursive when it calls itself from inside the body.
 Simple syntax is func r(){ r() }. If we run this function it will loop infinitely. Hence we use a if- else
 statement to come of the inifinite loop. A simple example of a recursive function is factorial of n.
*/
func getFactorial(num int) int {
  if num > 1 {
    return num * getFactorial(num - 1)
  }
  return 1
}

func main() {
  f := getFactorial(4)
  fmt.Println(f)
}
