// Converting Integer and Float Values to String

package main

import (
  "fmt"
  "strconv"
)

func main(){
    //number := int64(100)
    //numberStr := strconv.FormatInt(number, 10) // Note: 10 is the base // Note: FormatInt accepts int64 type
    //fmt.Println(numberStr)

    number := 100
    numberStr := strconv.Itoa(number)
    fmt.Println(numberStr)

    numberFloat := 23445221.12233
    numberFloatStr := strconv.FormatFloat(numberFloat, 'f', 3, 64) // Note: f is the format, 3 is the precision(You can use -1 if
      //you do not know the number of digits after the dot), 64 is the bitSize
    fmt.Println(numberFloatStr)
}
