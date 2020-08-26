// Extracting Unique elements from a list

package main

import (
  "fmt"
)

func main(){
  intSlice := []int{1,5,5,5,5,7,8,6,6, 6} // The int specifies the type of data in the list
  fmt.Println(intSlice)

  uniqueIntSlice := unique(intSlice)
  fmt.Println(uniqueIntSlice)
}

func unique(intSlice []int) []int{ // The second []int is the return value
    keys := make(map[int]bool)
    uniqueElements := []int{}

    for _,entry := range intSlice {
      if _, value := keys[entry]; !value{
        keys[entry] =true
        uniqueElements = append(uniqueElements, entry)
      }
    }
    return uniqueElements
}
