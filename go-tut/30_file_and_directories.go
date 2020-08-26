// Files and Directories

package main

import (
  "fmt"
  "os"
)

func main(){
    if _, err := os.Stat("log.txt"); os.IsNotExist(err){
        fmt.Println("Log.txt file does not exist")
    }else{
        fmt.Println("Log.txt file exists")
    }
}
