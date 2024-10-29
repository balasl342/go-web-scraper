// scraper/write.go
package scraper

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/balasl342/go-web-scraper/models"
)

// WriteToCSV writes the scraped articles to a CSV file.
func WriteToCSV(articles []models.Article, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Title", "URL"})

	// Write article rows
	for _, article := range articles {
		writer.Write([]string{article.Title, article.URL})
	}
}
