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

		go fetch(url, ch) // starts goroutine along with channel ch
	}

	file, _ := os.OpenFile("fetches.txt", os.O_CREATE|os.O_APPEND, 0666)

	for range os.Args[1:] {
		fmt.Fprintln(file, <-ch)
	}

	file.Close()

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed.\n", secs)
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
		ch <- fmt.Sprintf("fetch: %v", err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetch: reading from %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %d %s", secs, nbytes, url)
}
