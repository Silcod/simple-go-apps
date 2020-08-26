// Reading a particular line in a file

package main

import (
  "fmt"
  "os"
  "bufio"
)

func main(){
    fmt.Println(ReadLine(3)) // 3 is the line number
}

func ReadLine(lineNumber int) string{ // int is the data type of the parameter; 
    file, _ := os.Open("names.txt")
    fileScanner := bufio.NewScanner(file)
    lineCount := 1
    for fileScanner.Scan(){
      if lineCount == lineNumber{
          return fileScanner.Text()
      }
      lineCount++
    }
    defer file.Close()
    return ""
}
