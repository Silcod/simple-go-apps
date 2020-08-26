// Testing for the presence of a key in a map

package main

import (
  "fmt"
)

func main(){
    nameAges := map[string]int{
      "Tarik": 32,
      "Micheal": 30,
      "Jon": 25,
    }

    value, exists := nameAges["Tarik"] // value is the variable, exists is a boolean value which is true if there is a value for Tarik, it doesn't have to be exists, it can be anything
    fmt.Println(value)
    fmt.Println(exists)


    nameAges1 := map[string]int{
      "Tarik": 32,
      "Micheal": 30,
      "Jon": 25,
      "Jessica": 20,
    }

    if value, exists := nameAges1["Jesica"]; exists{
        fmt.Println(value)
      }else{
          fmt.Println("Jessica cannot be found")
      }
}
