// Creating a simple Web server

package main

import (
  "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello, World"))
}

func sayHelloMars(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello, Mars"))
}

func main(){
    http.HandleFunc("/", sayHello)
    http.HandleFunc("/mars", sayHelloMars)
    err := http.ListenAndServe(":5050", nil)
    if(err != nil){
      panic(err)
    }
}
