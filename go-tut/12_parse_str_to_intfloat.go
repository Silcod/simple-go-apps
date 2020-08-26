// Parsing String Values to Integer and Float

package main

import (
  "fmt"
  "strconv"
)

func main(){
  number := "2"
  valueInt, err := strconv.ParseInt(number, 10, 64) // 10 is the base, 64 is the bitSize
    if(err != nil){
      fmt.Print("Error happened.")
    }else{
      if (valueInt == 2){
        fmt.Println("Success")
      }
    }

    numberFloat := "2.2"
    valueFloat, errFloat := strconv.ParseFloat(numberFloat, 64) // 64 is the bitSize
      if(errFloat != nil){
        fmt.Print("Error happened.")
      }else{
        if (valueFloat == 2.2){
          fmt.Println("Success")
        }
      }
}
