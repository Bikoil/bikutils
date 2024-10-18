package main

import (
  "fmt"
  "os"
  "log"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: rwcheck <filename>")
    return
  }

  filename := os.Args[1]

  fileInfo, err := os.Stat(filename)
  if err != nil {
    log.Fatal(err)
  }

  mode := fileInfo.Mode()

  fmt.Println("Permissions:", mode)
  fmt.Printf("Numeric Permissions:", mode.Perm())
}
