package main

import (
  "fmt"
)
/*
  defer: is a keyword in Go that makes a function executes at the end of the execution of parent function or
   when parent function hits return statement
*/
func sayDone() {
  fmt.Println("I am done")
}

func main() {
  fmt.Println("main started")

  defer sayDone()

  fmt.Println("main finished")
}
