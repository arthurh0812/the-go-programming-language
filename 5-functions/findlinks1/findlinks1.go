package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	htmlprinter "github.com/arthurh0812/the-go-programming-language/5-functions/exercise-5.7"

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
	p := htmlprinter.NewPrinter()

	nodes := make(chan *html.Node)
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		go func(link string) {
			node, err := p.FetchDoc(link)
			if err != nil {
				log.Fatalf("findlinks1: %v\n", err)
			}
			nodes <- node
			wg.Done()
		}(link)
	}

	go func() {
		wg.Wait()
		close(nodes)
	}()

	for n := range nodes {
		p.PrintDoc(n)
	}
	log.Println("findlinks1: finished successfully")

}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" && strings.HasPrefix(attr.Val, "https://") {
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
