package main

import (
  "fmt"
)
/*
  defer: is a keyword in Go that makes a function executes at the end of the execution of parent function or
   when parent function hits return statement
*/
func endTime(timestamp string) {
  fmt.Println("Program ended at", timestamp)
}

func main() {
  time := "1 PM"

  defer endTime(time)

  time = "2 PM"

  fmt.Println("doing something")
  fmt.Println("main finished")
  fmt.Println("time is", time)
}

/*
  In the above program, we deferred execution of endTime function which means it will get executed at the
   end of main function but since at the end main function, time === "2 PM", we were expecting Program
    ended at 2 PM message. Even though because of deferred execution, endTime function is executing at
    the end of main function, it was pushed into the stack with all available argument values earlier
    when time variable was still 1 PM.
*/
