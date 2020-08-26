package main

import (
  "fmt"
  "time"
)
// Using goroutines is a great way to turn what will be a sequential program to a concurrent program
func compute(value int) {
  for i := 0; i < value; i++ {
    time.Sleep(time.Second)
    fmt.Println(i)
  }
}


func main() {
  fmt.Println("Concurrency With Goroutines")

  /*
    compute(5)
    compute(5) Here these two functions execute Synchronously, and it takes more time. To make the functions to run Asynchronously, add "go" in
    front of the functions and the functions will execute in less time. 
  */

  go compute(5)
  go compute(5)

  fmt.Scanln()
}
