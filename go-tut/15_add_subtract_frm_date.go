// Adding or Subtracting from Date

package main

import (
  "fmt"
  "time"
)

func main(){
    current := time.Now()
    febDate := current.AddDate(1, 1, 0) // years: + 1 year, months: + 1 month, days: 0

    fmt.Println(current.String())
    fmt.Println(febDate.String())

    // febDate.Sub(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))

    oneLessYears := febDate.AddDate(-1, 0, 0) // years: - 1 year, months: - 0 month, days: - 0 days
    fmt.Println(oneLessYears.String())

    tenMoreMinutes := febDate.Add(10 * time.Minute) // You can also use time.Hour
    fmt.Println(tenMoreMinutes)
}
