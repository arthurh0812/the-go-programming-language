// fetch3 prints the response body of the specified URL(s).
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		url = formatURL(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		n, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading from %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("read %d bytes from %s.\n", n, url)
	}
}

func formatURL(url string) string {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		return strings.Replace(url, "", "https://", 1)
	}
	return strings.Replace(url, "http://", "https://", 1)
}
