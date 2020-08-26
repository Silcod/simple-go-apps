// Processing Command Line Arguments

package main

import (
  "fmt"
  "os"
)

func main(){
    realArgs := os.Args[1:]

    if len(realArgs) == 0{
      fmt.Println("Please pass an argument.")
      return
    }

    fmt.Println(realArgs)
    if (realArgs[0] == "a") {
      writeHelloWorld()
    }else if realArgs[0] == "b"{
      writeHelloMars()
    }else{
      fmt.Println("Please pass a valid Arguments.")
    }
}
func writeHelloWorld(){
  fmt.Println("Hello, World")
}
func writeHelloMars(){
  fmt.Println("Hello, Mars")
}
