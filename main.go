package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Initialize the scraper and URL list
	scraper := NewScraper()

	// Visit the list of characters page and gather URLs
	scraper.GetCharacterURLs(&wg)

	// Start scraping each character
	scraper.ScrapeCharacters(&wg)

	// Wait for all scraping goroutines to finish
	go func() {
		wg.Wait()
		close(scraper.DataChannel)
	}()

	err := scraper.WriteDataToCSV("characters.csv")
	if err != nil {
		fmt.Println("Error writing data to CSV:", err)
	}

	fmt.Printf("Scraped %d characters\n", len(scraper.CharacterURLs))
}
