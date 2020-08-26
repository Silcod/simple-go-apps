// Replacing Parts of a String
package main

import (
  "fmt"
  "strings"
)

func main(){
  helloWorld := "Hello, World. How are you World, I am good, thanks World."
  helloMars := strings.Replace(helloWorld, "World", "Mars", 2) // If you don't want to count how many (world) there are, just use -1
  helloMars_all := strings.Replace(helloWorld, "World", "Mars", -1) 
  fmt.Println(helloMars)
  fmt.Println(helloMars_all)
}
