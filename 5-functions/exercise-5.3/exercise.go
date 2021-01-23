package main

import (
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	file, err := os.OpenFile("./texts.html", os.O_RDONLY, 500)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(file)
	file.Close()
	if err != nil {
		log.Fatalf("exercise-5.2: %v\n", err)
	}

	textPrint(os.Stdout, doc)
}

func printText(w io.Writer, n *html.Node) {
	n.Data = strings.Trim(n.Data, "\t\n\r\v\f ")
	w.Write([]byte(n.Data))
}

func textPrint(w io.Writer, n *html.Node) {
	if n.Data == "script" || n.Data == "style" {
		return
	}

	if n.Type == html.TextNode {
		printText(w, n)
	}

	for curr := n.FirstChild; curr != nil; curr = curr.NextSibling {
		textPrint(w, curr)
	}
}
