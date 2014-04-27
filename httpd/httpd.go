package main

import (
  "fmt"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello, world</h1>")
}

func main() {
  http.HandleFunc("/", hello)
  http.ListenAndServe(":12345", nil)
}
