// Writing to a file

package main

import (
  "fmt"
  "io/ioutil"
)

func main(){
    hello := "Hello, World Ibrahim"
    err := ioutil.WriteFile("hello_world", []byte(hello), 0644) // hello_world is the filename,[]byte(hello) is the data, 0644 is the file permissions
    if err != nil{
      fmt.Println(err)
    }
}
