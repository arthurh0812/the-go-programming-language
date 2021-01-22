package htmlprinter

import (
	"io"
	"log"
	"os"
	"sync"
)

// Default settings
var (
	// DefaultIndent is the indent that is used if none was specified
	DefaultIndent = "  "
	// DefaultOutput is the output writer that will be written to if none was specified
	DefaultOutput = os.Stdout
	// DefaultURL is the URL that the document will be fetched from if no URL was specified
	DefaultURL = "https://www.google.com"
)

// Printer is the actual pretty printer
type Printer struct {
	output io.Writer
	url    string
	indent string
	ending bool
	// shortcuts bool
	mu    *sync.Mutex
	depth int
}

// NewPrinter initiates a new Printer to pretty-print HTML
func NewPrinter() *Printer {
	return &Printer{
		indent: DefaultIndent,
		output: DefaultOutput,
		url:    DefaultURL,
		ending: true,
		// shortcuts: true,
	}
}

// URL specifies the URL that the document will be fetched from
func (p *Printer) URL(url string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.url = url
}

// WithIndent sets the indentation to print along for each
func (p *Printer) WithIndent(s string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.indent = s
}

// EndingTag decides whether to also print each ending tag or not
func (p *Printer) EndingTag(b bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.ending = b
}

// // Shortcuts decides whether to combine the starting and ending tag for some elements (e.g. <img />)
// func (p *Printer) Shortcuts(b bool) {
// 	p.mu.Lock()
// 	defer p.mu.Unlock()
// 	p.shortcuts = b
// }

// Output registers the output to which the HTML is going to be written
func (p *Printer) Output(w io.Writer) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.output = w
}

// Print sens an HTTP request to the provided URL. If none is provided it uses the printer's URL.
// Print prints out each HTML element and its attributes coming of the HTTP response.
func (p *Printer) Print(url string) {
	doc, err := p.fetchDoc(url)
	if err != nil {
		log.Fatal("html printer:", err.Error())
	}

	err = p.printDoc(doc)
	if err != nil {
		log.Fatal("html printer:", err.Error())
	}
}
