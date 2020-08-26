// Parsing String Values to Boolean

package main

import (
  "fmt"
  "strconv"
)

func main(){
  isNew := "true" // In Boolean: 1 means true and 0 means false, 
  isNewBool, err := strconv.ParseBool(isNew)
  if(err != nil){
    fmt.Println("failed")
  }else {
    if(isNewBool){
      fmt.Print("isNew")
    }else{
      fmt.Println("Not new")
    }
  }
}
