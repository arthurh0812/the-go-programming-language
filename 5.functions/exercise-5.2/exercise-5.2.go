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
		log.Fatalf("exercise-5.2: %v\n", err)
	}

	tagsCount := make(map[string]int)

	countTags(tagsCount, doc)

	for tag, count := range tagsCount {
		fmt.Println(tag, count)
	}
}

func addNodeTag(tags map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		tags[n.Data]++
	}
}

func countTags(tags map[string]int, n *html.Node) {
	addNodeTag(tags, n)

	if n.FirstChild != nil {
		countTags(tags, n.FirstChild)
	}

	if n.NextSibling != nil {
		countTags(tags, n.NextSibling)
	}
}
