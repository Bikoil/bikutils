package main

import (
  "fmt"
  "os"
  "log"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: rm <filename>")
    return
  }

  file := os.Args[1]
  err := os.Remove(file)
  if err != nil {
    log.Fatal(err)
  }
}
