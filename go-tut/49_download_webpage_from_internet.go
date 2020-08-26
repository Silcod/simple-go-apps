// Download a Web Page from Internet

package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main(){
    url := "http://golang.org"
    response, err := http.Get(url)
    if err != nil{
      panic(err)
    }

    defer response.Body.Close()

    html, err := ioutil.ReadAll(response.Body)
    if err != nil{
      panic(err)
    }

    fmt.Println(string(html))
}
