// Waiting for all concurrent functions to finish

package main

import (
  "fmt"
  "sync"
)

func main(){
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1) // 1 is delta
        go func(){
          fmt.Println("Hello World")
          wg.Done()
        }()
    }
    wg.Wait()
}
