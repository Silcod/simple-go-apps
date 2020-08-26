package main
// Hello, Concurrent world: the code is quite simple - single channel, single goroutine, one write, one read
func main() {
  // create the new channel of type int
  ch := make(chan int)

  // start new anonymous goroutine
  go func(){
    // send 42 to channel
    ch <- 42
  }()
  // read from channel
  <- ch
}
// To visualize go to this link: https://divan.dev/demos/hello/
