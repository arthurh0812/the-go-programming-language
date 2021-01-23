package main

import (
	"fmt"
	"log"
	"os"

	htmlprinter "github.com/Arthur0812/private-go/the-go-programming-language/5.functions/exercise-5.7"

	"golang.org/x/net/html"
)

func main() {
	htmlFile, err := os.OpenFile("./links.html", os.O_RDONLY, 700)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(htmlFile)
	if err != nil {
		log.Fatalf("findlinks1: %v\n", err)
	}

	links := []string{}

	links = visit(links, doc)

	for _, link := range links {
		fmt.Println(link)
	}

	p := htmlprinter.NewPrinter()
	p.Print("")
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	// recursive call of visit for each of the siblings inside n
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func changeSlice(sl []string) {
	// accesses the underlying array by index
	// => AFFECTS ARGUMENT
	sl[0], sl[2] = "3", "4"
	fmt.Println("len:", len(sl))
	fmt.Println("cap:", cap(sl))
	fmt.Println("sl:", sl)

	// resets the argument slice reference
	// => NO AFFECT ON ARGUMENT
	sl = sl[:1]
	fmt.Println("len:", len(sl))
	fmt.Println("cap:", cap(sl))
	fmt.Println("sl:", sl)

	// resets the argument slice reference to a newly created one
	// => NO AFFECT ON ARGUMENT
	sl = []string{"1", "2"}
	fmt.Println("len:", len(sl))
	fmt.Println("cap:", cap(sl))
	fmt.Println("sl:", sl)

	// append returns a reference to a new slice
	// => NO AFFECT ON ARGUMENT
	sl = append(sl, "10")
	fmt.Println("len:", len(sl))
	fmt.Println("cap:", cap(sl))
	fmt.Println("sl:", sl)
}
