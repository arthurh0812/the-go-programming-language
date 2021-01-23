package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	var all []string
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("exercise-5.4: fetching document: url '%v' not found: %v", url, err)
		}

		if cd := resp.StatusCode; cd < 200 || cd >= 400 {
			log.Fatalf("exercise-5.4: reading response status: error status code: %v", cd)
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatalf("exercise-5.4: reading response body: fail: %v", err)
		}

		all = append(all, findLinks(doc)...)
	}

	fmt.Printf("Found %d links:\n", len(all))
	for _, link := range all {
		fmt.Printf("%v\n", link)
	}
}

func visit(n *html.Node) []string {
	links := make([]string, 0)
	if n.Type == html.ElementNode {
		if n.Data == "a" || n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		} else if n.Data == "img" || n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	return links
}

func findLinks(node *html.Node) (links []string) {
	l := visit(node)

	for curr := node.FirstChild; curr != nil; curr = curr.NextSibling {
		links = append(links, findLinks(curr)...)
	}

	return append(links, l...)
}
