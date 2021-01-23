package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, img, err := CountWordsAndImages(url)
		if err != nil {
			log.Fatalf("exercise-5.5: %v", err)
		}
		fmt.Printf("Words:\t%d\n", words)
		fmt.Printf("Images:\t%d\n", img)
	}
}

// CountWordsAndImages sends an HTTP GET request to the given url.
// CountWordsAndImages counts the words and images of the response
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("sending request: fail: %v", err)
	}
	if cd := resp.StatusCode; cd < 200 || cd >= 400 {
		return 0, 0, fmt.Errorf("reading response: error code: %v, msg: %v", cd, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0, 0, fmt.Errorf("parsing body: fail: %v", err)
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		buf := bytes.NewBuffer([]byte(n.Data))
		scanner := bufio.NewScanner(buf)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	} else if n.Data == "img" {
		images++
	}

	for curr := n.FirstChild; curr != nil; curr = curr.NextSibling {
		w, i := countWordsAndImages(curr)
		words, images = words+w, images+i
	}

	return
}
