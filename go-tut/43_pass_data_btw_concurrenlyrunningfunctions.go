// Passing Data Between Concurrently Running Functions

package main

import (
  "fmt"
  //"time"
)

func main(){
    nameChannel := make(chan string)
    done := make(chan string)
    go func ()  {
      names := []string{"tarik", "micheal", "gopi", "jessica"}
      for _, name := range names {
        // doing some operation
        fmt.Println("Processing the first stage of: " + name)
        nameChannel <- name // writing to the channel nameChannel
      }
      close(nameChannel)
    }()

    go func ()  {
      for name := range nameChannel{
          fmt.Println("Processing the second stage of: " + name)
      }
      done <- "" // To close a channel
    }()
    <-done // just know the program won't work without this
   // time.Sleep(1 * time.Second) // Or this
}
