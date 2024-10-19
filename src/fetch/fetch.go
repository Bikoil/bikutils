package main

import (
  "io"
  "net/http"
  "os"
  "log"
  "path"
  "net/url"
)

func main() {
  var fileURL string
  
  if len(os.Args) > 1 {
    fileURL = os.Args[1]
  } else {
      log.Println("Usage: fetch <download link>")
    return
  }

  parsedURL, err := url.Parse(fileURL)
  if err != nil {
    log.Fatal(err)
  }

  filename := path.Base(parsedURL.Path)

  resp, err := http.Get(fileURL)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  out, err := os.Create(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer out.Close()

  _, err = io.Copy(out, resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  log.Println("File fetched successfully!", filename)
}
