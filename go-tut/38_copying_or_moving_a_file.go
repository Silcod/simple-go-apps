// Copying or Moving a file

package main

import (
  "os"
  "io"
)

func main(){
    original, err := os.Open("original.txt")

    if err != nil{
      panic(err)
    }

    defer original.Close()

    original_copy, err2 := os.Create("copy.txt")

    if err2 != nil{
      panic(err2)
    }

    defer original_copy.Close()

    _, err3 := io.Copy(original_copy, original)

    if err3 != nil{
      panic(err3)
    }
}
