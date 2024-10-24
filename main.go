// main.go
package main

import (
	"fmt"
	"log"

	"github.com/balasl342/go-web-scraper/scraper"
)

func main() {
	// URL to scrape
	baseURL := "https://news.ycombinator.com/"
	fmt.Printf("Scraping URL: %s\n", baseURL)

	// Fetch the HTML content
	doc, err := scraper.FetchHTML(baseURL)
	if err != nil {
		log.Fatalf("Error fetching HTML: %v", err)
	}
	fmt.Printf("Fetched HTML content%+v", doc)
}
