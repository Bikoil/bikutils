package main

import (
  "fmt"
  "os"
  "log"
)

func main() {
  if len(os.Args) > 1 {
    dir := os.Args[1]
    files, err := os.ReadDir(dir)
    if err != nil { 
      log.Fatal(err)
    }

    for _, file := range files { 
      fmt.Println(file.Name())
    }
  } else {
      files, err := os.ReadDir(".")
      if err != nil { 
        log.Fatal(err)
      }

      for _, file := range files { 
        fmt.Println(file.Name())
      }
   }
}
