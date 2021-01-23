package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// NodeFunc is simply a traversal function used on a node
type NodeFunc func(n *html.Node) bool

func main() {
	for i := 1; i < len(os.Args); i += 2 {
		resp, err := http.Get(os.Args[i])
		if err != nil {
			log.Fatalf("sending request: fail: %v", err)
		}

		if cd := resp.StatusCode; cd < 200 || cd >= 400 {
			log.Fatalf("reading response: error status code: %v", cd)
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatalf("parsing to html: fail: %v", err)
		}

		n := GetElementByID(doc, os.Args[i+1])

		if n == nil {
			fmt.Println("no such element found...")
			return
		}
		fmt.Printf("%#v\n", n)
	}
}

// GetElementByID returns the first HTML element with the given ID as an attribute
func GetElementByID(doc *html.Node, id string) *html.Node {
	hasID := func(n *html.Node) bool {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
		return false
	}

	return forEachNode(doc, hasID, hasID)
}

func forEachNode(n *html.Node, pre, post NodeFunc) *html.Node {
	if pre != nil && n != nil {
		if pre(n) {
			return n
		}
	}

	for curr := n.FirstChild; curr != nil; curr = curr.NextSibling {
		n := forEachNode(curr, pre, post)
		if n != nil {
			return n
		}
	}

	if post != nil {
		if post(n) {
			return n
		}
	}

	return nil
}
