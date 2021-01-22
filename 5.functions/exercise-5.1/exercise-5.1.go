package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	file, err := os.OpenFile("./links.html", os.O_RDONLY, 500)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalf("exercise-5.1: %v\n", err)
	}

	links := visit([]string{}, doc)
	for _, link := range links {
		fmt.Println(link)
	}
}

func getHref(n *html.Node) (string, bool) {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				return attr.Val, true
			}
		}
	}
	return "", false
}

func visit(links []string, n *html.Node) []string {
	// get "href" attribute of n (if existing)
	href, ok := getHref(n)
	if ok {
		links = append(links, href)
	}

	// execute visit on n's first child (one layer deeper)
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	// and on n's next sibling (same layer, node on the right)
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}
