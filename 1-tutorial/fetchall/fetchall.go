// fetchall fetches all specified URLs concurrently and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		url = formatURL(url)
		go fetch(url, ch) // start a goroutine along with channel ch
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed.\n", time.Since(start).Seconds())
}

func formatURL(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return strings.Replace(url, "", "https://", 1)
	}
	return strings.Replace(url, "http://", "https://", 1)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v", err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("fetch: reading from %s: %v", url, err) // send to channel ch
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %d %s", secs, nbytes, url) // send to channel ch
}
