// Deleting a directory and its content

package main

import (
  "fmt"
  "os"
)

func main(){
    err := os.Remove("hello") // To remove both directory and its contents us "os.RemoveAll"
    if err != nil{
      fmt.Println(err)
    }
}
