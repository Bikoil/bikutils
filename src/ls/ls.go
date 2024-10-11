package main
import (
  "fmt"
  "os"
)

func main() {
  files, _ := os.ReadDir(".")
  for _, file := range files {
    fmt.Println(file.Name())
  }
}
