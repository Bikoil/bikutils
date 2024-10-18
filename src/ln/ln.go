package main

import (
  "flag"
  "fmt"
  "os"
)

func main() {
  sFlag := flag.Bool("s", false, "create a symbolic link")
  flag.Parse()

  if flag.NArg() != 2 {
    fmt.Println("Usage: ln [-s] <target> <link>")
    return
  }

  target := flag.Arg(0)
  link := flag.Arg(1)

  if *sFlag {

    err := os.Symlink(target, link)
    if err != nil {
      fmt.Println("Error creating symlink:", err)
      return
    }
    fmt.Println("Symbolic link created:", link, "->", target)
  } else {
    
    err := os.Link(target, link)
    if err != nil {
      fmt.Println("Error creating hard link:", err)
      return
    }
    fmt.Println("Hard link created:", link, "->", target)
  }
}
