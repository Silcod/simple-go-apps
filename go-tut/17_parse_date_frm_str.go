// Parsing Dates and Times from strings

package main

import (
  "fmt"
  "time"
)

func main(){
    str := "2020-01-09T11:45:26.371Z"
    layout := "2006-01-02T15:04:05.000Z"

    t, err := time.Parse(layout, str)

    if err != nil{
      fmt.Println(err)
    }

    fmt.Println(t.String())
}
