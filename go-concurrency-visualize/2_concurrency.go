package main

import (
  "fmt"
  "time"
)
// Timers
func timer(d time.Duration) <-chan int {
  c := make(chan int)
  go func() {
    time.Sleep(d)
    c <- 1
  }()
  return c
}

func main() {
  for i := 0; i < 24; i++ {
    c := timer(1 * time.Second)
    <-c
    fmt.Println(i) // prints 1 to 23
  }
}
// To visualize go to this link: https://divan.dev/demos/timers/
/*
  In fact, you can build a simple timer with this approach - create a channel, start goroutine which
   writes to this channel after given duration and returns this channel to the caller of your func.
   The caller then blocks on reading from the channel for the exact amount of time. Let’s run such
   timer 24 times and try to visualize it.

*/
