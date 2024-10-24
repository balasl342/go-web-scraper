// scraper/fetch.go
package scraper

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// FetchHTML fetches the HTML content of the given URL and returns a parsed HTML node.
func FetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	// Parse the HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	return doc, nil
}
