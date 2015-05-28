package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", Index)

  log.Fatal(http.ListenAndServe(":5000", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
