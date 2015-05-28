package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

const DefaultPort = ":3000"
const EnvPort = "PORT"

func main() {
	http.HandleFunc("/", Index)

	port := DefaultPort
	if os.Getenv(EnvPort) != "" {
		port = ":" + os.Getenv(EnvPort)
	}

	log.Fatal(http.ListenAndServe(port, nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
