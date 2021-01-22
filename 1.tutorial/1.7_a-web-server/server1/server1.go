// server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL = %q\n", req.URL)
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
