// Iterating over an array

package main

import (
  "fmt"
)

func main(){
    numbers := []int{1, 5, 3, 6, 2, 10, 8}

    for index,value := range numbers{
        fmt.Printf("Index: %v and Value: %v\n", index, value)
    }

    // If you are not interested in the index anymore
    numberstwo := []int{1, 5, 3, 6, 2, 10, 8}

    for _,value := range numberstwo{
        fmt.Println(value)
    }
}
