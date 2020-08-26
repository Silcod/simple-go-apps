/*
  Workers: The opposite pattern to fan-in is a fan-out or workers pattern. Multiple goroutines can
  read from a single channel, distributing an amount of work between CPU cores, hence the workers name.
   In Go, this pattern is easy to implement - just start a number of goroutines with channel as parameter,
    and just send values to that channel - distributing and multiplexing will be done by Go runtime,
     automagically :)
*/


// To visualize it, go to this link: https://divan.dev/demos/workers/
package main

import (
  "fmt"
  "time"
  "sync"
)

func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
  defer wg.Done()
  for {
    task, ok := <-tasksCh
    if !ok {
      return
    }

    d := time.Duration(task) * time.Millisecond
    time.Sleep(d)
    fmt.Println("processing task", task)
  }
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
  tasksCh := make(chan int)

  for i := 0; i < workers; i++ {
    go worker(tasksCh, wg)
  }

  for i := 0; i < tasks; i++ {
    tasksCh <- i
  }
  close(tasksCh)
}

func main() {
  var wg sync.WaitGroup
  wg.Add(36)
  go pool(&wg, 36, 50)
  wg.Wait()
}
