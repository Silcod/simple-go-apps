// Running multiple functions concurrently

package main

import "time"
import "fmt"

func main(){
    go func ()  { // with the keyword go, you are telling golang to run this particular function as a go routine
      names := []string{"tarik", "john", "micheal", "jessica"}
      for _, name := range names {
        time.Sleep(1 * time.Second)
        fmt.Println(name)
      }
    }()

    go func ()  {
      ages := []int{1, 2, 3, 4, 5}

      for _, age := range ages {
        time.Sleep(1 * time.Second)
        fmt.Println(age)
      }
    }()
    time.Sleep(3 * time.Second) // this is the overall time you are giving the two functions to run, so if you give it enough time, it will finish running
}

/*func main() {
    nameChannel := make(chan string) // To define a channel

    go func ()  { // with the keyword go, you are telling golang to run this particular function as a go routine
      names := []string{"tarik", "john", "micheal", "jessica"}

      for _, name := range names {
        time.Sleep(1 * time.Second)
        fmt.Println(name)
      }
      nameChannel <- "" // To close a channel
    }()
    <-nameChannel
}
*/