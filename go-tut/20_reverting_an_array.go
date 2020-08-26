// Reverting an Array

package main

import (
  "fmt"
  "sort"
)

func main(){
    // To sort in Ascending order
    numbers := []int{1, 5, 3, 6, 2, 10, 8}
    tobeSorted := sort.IntSlice(numbers) // sort.IntSlice() function will give you: [1 5 3 6 2 10 8]
    sort.Sort(tobeSorted)
    fmt.Println(tobeSorted)

    // To sort in Descending order
    numbers_asc := []int{1, 5, 3, 6, 2, 10, 8}
    tobeSorted_asc := sort.IntSlice(numbers_asc) // sort.IntSlice() function will give you: [1 5 3 6 2 10 8]
    sort.Sort(sort.Reverse(tobeSorted_asc))
    fmt.Println(tobeSorted_asc)

}
