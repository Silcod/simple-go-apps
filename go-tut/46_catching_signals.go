// Catching Signals
// Signals are limited form of interprocess communications and typically used in linux and linux like operating systems
// The signal is an async notification sent to a specific threat in the same process or another target process to notify it of an event occurrence

package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
)

func main(){
    signals := make(chan os.Signal, 1)
    done := make(chan bool)

    signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

    go func () {
        sig := <- signals
        fmt.Println(sig)
        fmt.Println("Signal captured and processed")
        done <- true
    }()

    fmt.Println("Waiting for signal")
    <-done
    fmt.Println("Exiting the application...")
}
