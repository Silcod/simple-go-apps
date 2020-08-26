package main
/*
  Ping Pong: This nice concurrency example was found in a great talk by googler Sameer Ajmani.
  Here we have a channel as a table of the ping-pong game. The ball is an integer variable, and two
  goroutines-players that ‘hit’ the ball, increasing its value (hits counter).
*/

import (
  "time"
  "fmt"
)
/*
  Go channels: when you have multiple goroutines which are executed concurrently, channels provide
  the easiest way to allow the goroutines to communicate with each other.
*/
func main() {
  var Ball int
  table := make(chan int)
  go player(table)
  go player(table)

  table <- Ball
  time.Sleep(1 * time.Second)
  <-table
  // fmt.Println(Ball) // If run prints 0
}

func player(table chan int) {
  for {
    ball := <-table
    ball++
    time.Sleep(100 * time.Millisecond)
    table <- ball
    //fmt.Println(ball) // If run print 1 to 10
  }
}
// To visualize it, go to this link: https://divan.dev/demos/pingpong/ and https://divan.dev/demos/pingpong3/
