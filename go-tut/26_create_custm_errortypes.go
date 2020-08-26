// Creating Custom Error Typespackage main

package main

import (
  "fmt"
)

type MyError struct {
  ShortMessage string
  DetailedMessage string
}

func (e *MyError) Error() string{ // e *MyError links e to MyError struct that means you can access MyError struct with e(eg: e.ShortMessage)
    // the string after Error() means that the function returns a string value back
    return e.ShortMessage + "\n" + e.DetailedMessage
}

func main(){
    err := doSomething()
    fmt.Println(err)
}

func doSomething() error{
    // Doing something here...
    return &MyError{ShortMessage:"Wohoo something happened!", DetailedMessage:"File cannot be found!"}
}
