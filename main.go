// main.go
package main

import (
	"fmt"
	"time"

	"github.com/balasl342/go-web-scraper/scraper"
	"github.com/balasl342/go-web-scraper/utils"
)

func main() {
	// URL to scrape
	baseURL := "https://www.hindustantimes.com/cricket"
	fmt.Printf("Scraping URL: %s\n", baseURL)

	// Fetch the HTML content
	doc, err := scraper.FetchHTML(baseURL)
	utils.CheckError("Error fetching HTML: ", err)

	// Extract articles from the HTML
	articles := scraper.ExtractArticles(doc)

	// Print and store the results
	for _, article := range articles {
		fmt.Printf("Title: %s\nURL: %s\n", article.Title, article.URL)
	}

	// Write to CSV file
	filename := fmt.Sprintf("articles_%d.csv", time.Now().Unix())
	scraper.WriteToCSV(articles, filename)
	fmt.Printf("Articles saved to %s\n", filename)
}
