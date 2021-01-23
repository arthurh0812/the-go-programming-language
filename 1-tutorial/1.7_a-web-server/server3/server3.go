// server3 echoes the HTTP request.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request
func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", req.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", req.RemoteAddr)
	if err := req.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range req.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
