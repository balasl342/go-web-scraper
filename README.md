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

## Prerequisites
- **Go 1.16+**
- Basic understanding of Golang

## Installation
1. **Clone the repository**:
   ```bash
   git clone https://github.com/balasl342/go-web-scraper.git
   ```

2. **Navigate to the project directory**:
   ```bash
   cd go-web-scraper
   ```

3. **Install dependencies** (if any) using Go modules:
   ```bash
   go mod tidy
   ```

## Usage
1. **Update the URL** in `main.go` (if needed). By default, it scrapes the Hindustan Times Cricket section.
   ```go
   baseURL := "https://www.hindustantimes.com/cricket"
   ```

2. **Run the scraper**:
   ```bash
   go run main.go
   ```

3. **Output**: The scraper prints the extracted titles and URLs to the console and stores them in a CSV file named something like `articles_<timestamp>.csv`.
