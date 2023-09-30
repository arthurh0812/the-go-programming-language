package htmlprinter

import (
	"fmt"

	"golang.org/x/net/html"
)

func (p *Printer) PrintDoc(doc *html.Node) error {
	if p.output == nil {
		return fmt.Errorf("printing document: no output provided")
	}

	preFunc := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Fprintf(p.output, "%*s<%s", p.depth, p.indent, n.Data)
			// add the attributes
			for i, attr := range n.Attr {
				if i == 0 {
					fmt.Fprint(p.output, " ")
				}
				fmt.Fprintf(p.output, "%s='%s'", attr.Key, attr.Val)
			}
			fmt.Fprint(p.output, ">\n")
			p.depth++
		}
	}

	postFunc := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Fprintf(p.output, "%*s</%s>\n", p.depth, p.indent, n.Data)
			p.depth--
		}
	}

	if !p.ending {
		postFunc = nil
	}

	forEachNode(doc, preFunc, postFunc)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for curr := n.FirstChild; curr != nil; curr = curr.NextSibling {
		forEachNode(curr, pre, post)
	}

	if post != nil {
		post(n)
	}
}
