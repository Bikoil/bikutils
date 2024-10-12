package main

import (
  "fmt"
  "os"
  "log"
)

func main() {
  dir, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(dir)
}
