// Converting a Byte Array to String

package main

import (
  "fmt"
)

func main(){
  // To convert byte array to string
  helloWorldByte := []byte {72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100} // The byte specifies the type of data in the list
  fmt.Println(string(helloWorldByte))

  // To convert string to Byte array
  helloWorld := "Hello, World"
  fmt.Println([]byte(helloWorld))
}
