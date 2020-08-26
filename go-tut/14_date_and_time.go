// Finding today's Date and Time

package main

import (
  "fmt"
  "time"
)

func main(){
  current := time.Now()
  fmt.Println(current.String())

  fmt.Println("MM-DD-YY : ", current.Format("01-09-2020"))

  fmt.Println("YY-MM-DD hh:mm:ss : ", current.Format("2020-01-09 16:02:53"))

}
