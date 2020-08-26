// Finding an Element from an Array

package main

import (
  "fmt"
  "sort"
)

func main(){
    str := []string{"St. George", "Provo", "Sandy", "Salt Lake City", "Draper", "South Jordan", "Murray"} // the string specifies the type of the data in the list

    for i, v := range str{ // Note: i is the index and v is the value
      if (v == "Sandy"){
          fmt.Println(i)
      }
    }

    sortedList := sort.StringSlice(str) // sort.StringSlice() will give you this: [St. George Provo Sandy Salt Lake City Draper South Jordan Murray]
    sortedList.Sort()
    fmt.Println(sortedList)

    index := sortedList.Search("Sandy")
    fmt.Println(index)
}
