// Doing a Simple Logging in your App

package main

import (
  "os"
  "fmt"
  "log"
)

func main(){
    log_file, err = os.Create("log_file")

    if err != nil{
      fmt.Println("An error occured...")
    }

    defer log_file.Close()

    log.SetOutput(log_file)

    log.Println("Doing some logging here...")
    log.Fatalln("Fatal: Application crashed!")
}
