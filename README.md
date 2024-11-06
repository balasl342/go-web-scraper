# Go Web Scraper

## Overview
This is a simple web scraper example built in Golang to extract news articles from a given website. In this example, the scraper is configured to fetch and parse article titles and URLs from the **Hindustan Times Cricket section**.

## Features
- Extracts article titles and URLs from a webpage
- Stores scraped data in a CSV file for easy analysis
- Handles relative and absolute URLs
- Modular code structure for easy maintenance

## Project Structure
```
go-web-scraper/
│
├── main.go                 # Main entry point for the program
├── scraper/
│   ├── fetch.go            # Functions to fetch and parse HTML content
│   ├── parse.go            # Functions to extract specific elements from the HTML content
│   └── write.go            # Functions to write scraped data to a CSV file
├── models/
│   └── article.go          # Defines the Article model
├── utils/
│   └── helper.go          # Utility functions like error handling
└── go.mod                  # Go module file
```
