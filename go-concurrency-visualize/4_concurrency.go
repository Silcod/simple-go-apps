/*
  Fan-In: To be short, fan-in is a function reading from the multiple inputs and multiplexing all into
  the single channel.
*/
/*
  Go channels: when you have multiple goroutines which are executed concurrently, channels provide
  the easiest way to allow the goroutines to communicate with each other.
*/

// To visualize it, go to this link: https://divan.dev/demos/fanin/
package main

import (
  "fmt"
  "time"
)

func producer(ch chan int, d time.Duration) {
  var i int
  for {
    ch <- i
    i++
    time.Sleep(d)
  }
}

func reader(out chan int) {
  for x := range out {
    fmt.Println(x)
  }
}

func main() {
  ch := make(chan int)
  out := make(chan int)
  go producer(ch, 100*time.Millisecond)
  go producer(ch, 250*time.Millisecond)
  go reader(out)
  for i := range ch {
    out <- i
    fmt.Println(i)
  }
}
