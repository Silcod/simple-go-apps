// Creating a simple file server

package main

import (
  "net/http"
)

func main(){
    http.Handle("/", http.FileServer(http.Dir("./images")))
    http.ListenAndServe(":5050", nil)

}
