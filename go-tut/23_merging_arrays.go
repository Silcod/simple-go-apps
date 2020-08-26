// Merging Arrays

package main

import (
  "fmt"
)

func main(){
  items1 := []int{3, 4}
  items2 := []int{1, 2}

  result := append(items1, items2...) // The ... is important
  fmt.Println(result)
}
