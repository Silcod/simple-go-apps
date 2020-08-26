// Counting lines in a file

package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
    file, _ := os.Open("names.txt")
    fileScanner := bufio.NewScanner(file)
    lineCount := 0;
    for fileScanner.Scan(){
      lineCount++
    }
    defer file.Close()
    fmt.Println(lineCount)
}
