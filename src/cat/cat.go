package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
)

func main() {
  nFlag := flag.Bool("n", false, "number all output lines")
  flag.Parse()

  if flag.NArg() == 0 {
    fmt.Println("Usage: cat [-n] <filename>")
    return
  }

  filename := flag.Arg(0)
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  lineNumber := 1
  for scanner.Scan() {
    line := scanner.Text()
    if *nFlag {
      fmt.Printf("%d %s\n", lineNumber, line)
      lineNumber++
    } else {
      fmt.Println(line)
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file:", err)
  }
}
