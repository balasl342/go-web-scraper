// scraper/parse.go
package scraper

import (
	"strings"

	"github.com/balasl342/go-web-scraper/models"
	"golang.org/x/net/html"
)

// ExtractArticles extracts articles from the given HTML document.
func ExtractArticles(doc *html.Node) []models.Article {
	var articles []models.Article
	var f func(*html.Node)

	// Recursive function to traverse the HTML node tree.
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// Look for <a> tags inside specific div classes like "headingfour"
			for _, attr := range n.Attr {
				if attr.Key == "href" && strings.Contains(attr.Val, "/cricket/") {
					// Extract title (inner text of the <a> tag)
					title := extractText(n)
					if title != "" {
						// Construct a full URL if the href attribute is a relative path
						url := attr.Val
						if !strings.HasPrefix(url, "https") {
							url = "https://www.hindustantimes.com" + url
						}

						articles = append(articles, models.Article{
							Title: title,
							URL:   url,
						})
					}
				}
			}
		}

		// Continue traversing the HTML tree
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	// Start traversing from the root node
	f(doc)
	return articles
}

// ExtractText recursively extracts the inner text of an HTML node.
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return strings.TrimSpace(text)
}
