package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "bytes"
)

func main() {
    response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD") // Using Coinbase API endpoint url
    if err != nil{
        fmt.Printf("The http request failed with error %s\n", err)
    }else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data)) // The data is going to be a slice of bytes so we have to convert to string with string() function
    }

    jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
    jsonValue, _ := json.Marshal(jsonData) // To encode JSON data we use the json.Marshal() function. To decode JSON data we use the Unmarshal function.
    response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
    if err != nil{
        fmt.Printf("The http request failed with error %s\n", err)
    }else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data)) // The data is going to be a slice of bytes so we have to convert to string with string() function
    }
}
