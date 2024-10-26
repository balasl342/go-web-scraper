// main.go
package main

import (
	"fmt"
	"log"

	"github.com/balasl342/go-web-scraper/scraper"
)

func main() {
	// URL to scrape
	baseURL := "https://www.hindustantimes.com/cricket"
	fmt.Printf("Scraping URL: %s\n", baseURL)

	// Fetch the HTML content
	doc, err := scraper.FetchHTML(baseURL)
	if err != nil {
		log.Fatalf("Error fetching HTML: %v", err)
	}

	// Extract articles from the HTML
	articles := scraper.ExtractArticles(doc)

	// Print and store the results
	for _, article := range articles {
		fmt.Printf("Title: %s\nURL: %s\n", article.Title, article.URL)
	}

}
