// Reading the entire content of a text file

package main

import (
  "fmt"
  "io/ioutil" // This package provides a read file function
)

func main(){
    contentBytes, err := ioutil.ReadFile("names.txt")
    if err == nil{
      var contentStr = string(contentBytes)
      fmt.Println(contentStr)
    }
}
