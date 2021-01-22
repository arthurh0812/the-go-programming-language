package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	file, err := os.OpenFile("./texts.html", os.O_RDONLY, 500)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalf("exercise-5.2: %v\n", err)
	}

	textPrint(os.Stdout, doc)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":5000", nil)
}

func printContent(w io.Writer, n *html.Node) {
	if n.Type == html.TextNode && n.Type != html.CommentNode {
		fmt.Fprintf(w, "Text Node!\n")
	}
	fmt.Fprintf(w, "n.Attr: %s\n", n.Attr)
	fmt.Fprintf(w, "n.Data: %s\n", n.Data)
	fmt.Fprintf(w, "n.DataAtom: %s\n", n.DataAtom)
	fmt.Fprintf(w, "n.Type: %d\n", n.Type)
	fmt.Fprintf(w, "n.Namespace: %s\n", n.Namespace)
	fmt.Fprintln(w, "----------------------------------")
}

func textPrint(w io.Writer, n *html.Node) {
	// don't call textPrint recursively on a script or style tag
	if n.Data == "script" || n.Data == "style" {
		return
	}

	printContent(w, n)

	if n.FirstChild != nil {
		textPrint(w, n.FirstChild)
	}
	if n.NextSibling != nil {
		textPrint(w, n.NextSibling)
	}
}
