// Converting Boolean to String

package main

import (
  "fmt"
  //"strconv"
)

func main(){
  isNew := true
  //isNewStr := strconv.FormatBool(isNew)
  message := "Purchase item is "
  fmt.Printf("%s %v", message, isNew)
}
