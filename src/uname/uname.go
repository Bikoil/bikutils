package main

import (
  "flag"
  "fmt"
  "os"
  "runtime"
)

func main() {
  aFlag := flag.Bool("a", false, "Display all system information")
  flag.Parse()

  if *aFlag {
    fmt.Println(runtime.GOOS, hostname(), "0.0", "1", runtime.GOARCH, runtime.Compiler)
  } else {
    fmt.Println(runtime.GOOS)
  }
}

func hostname() string {
  name, err := os.Hostname()
  if err != nil {
    return "unknown"
  }
  return name
}
