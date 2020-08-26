// List all the Files under a Directory

package main

import (
  "fmt"
  "io/ioutil"
)

func main(){
    files, err := ioutil.ReadDir("hello")

    if err != nil{
      panic(nil)
    }

    for _, f := range files{
      fmt.Println(f.Name())
    }
}
