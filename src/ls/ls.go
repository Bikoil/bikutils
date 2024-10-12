package main

import (
    "fmt"
    "os"
)

func main() {
  if len(os.Args) > 1 {
        
    dir := os.Args[1]
    files, _ := os.ReadDir(dir)
    for _, file := range files {
      fmt.Println(file.Name())
    }
  } else {
      files, _ := os.ReadDir(".")
      for _, file := range files {
        fmt.Println(file.Name())
      }
    }
}
