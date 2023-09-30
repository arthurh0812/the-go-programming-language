package htmlprinter

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func (p *Printer) FetchDoc(url string) (*html.Node, error) {
	if url == "" {
		if p.url == "" {
			return nil, fmt.Errorf("using url: no vlid url can be used")
		}
		url = p.url
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching document: http request failed for %v", url)
	}

	if code := response.StatusCode; !statusSuccessful(code) {
		return nil, fmt.Errorf("reading response: http status code unsuccessful: %v", code)
	}

	doc, err := html.Parse(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing document: html parsing failed: %v", err)
	}

	return doc, nil
}

func statusSuccessful(code int) bool {
	return code >= 200 && code < 400
}
