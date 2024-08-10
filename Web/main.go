package main

import (
  "fmt"
  "net/http"
)

func main() {
  fmt.Println("Welcome to Dummy Bank Web!")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Dummy Bank Web!")
  })
  http.ListenAndServe(":9000", nil)
}
