// Finding Difference btw Two Dates

package main

import (
  "fmt"
  "time"
)

func main(){
    first := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC) // year: 200, month: 1, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, time-zone: time.UTC
    second := time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC)

    difference := second.Sub(first)
    fmt.Printf("Difference %v", difference)
}
