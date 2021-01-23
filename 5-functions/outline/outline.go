package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	htmlFile, err := os.OpenFile("./links.html", os.O_RDONLY, 500)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(htmlFile)
	if err != nil {
		log.Fatalf("outline: %v\n", err)
	}

	outline([]string{}, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // add the node's HTML tag
		fmt.Println(stack)
	}

	for curr := n.FirstChild; curr != nil; curr = curr.NextSibling {
		outline(stack, curr)
	}
}
