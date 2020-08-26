package main

import (
  "fmt"
)

func greet(message string) {
  fmt.Println("greeting: ", message)
}

func main() {
  fmt.Println("Call one")

  defer greet("Greet one")

  fmt.Println("Call two")

  defer greet("Greet two")

  fmt.Println("Call three")

  defer greet("Greet three")
}
/*
  Practical use of defer can be seen when a function has too many conditions whether if-else or case
   statements, and at the end of each condition, you need to do something like close a file or send
   http response. Instead of writing multiple calls, we can use defer to save the day.
*/
