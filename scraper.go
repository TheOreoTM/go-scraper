package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"github.com/gocolly/colly"
)

type Scraper struct {
	CharacterURLs []string
	URLSet        map[string]struct{}
	DataChannel   chan map[string]string
	ShouldDebug   bool
}

// NewScraper initializes a new Scraper
func NewScraper() *Scraper {
	return &Scraper{
		CharacterURLs: []string{},
		URLSet:        make(map[string]struct{}),
		DataChannel:   make(chan map[string]string),
		ShouldDebug:   false,
	}
}

// GetCharacterURLs gathers all unique character URLs from the list page
func (s *Scraper) GetCharacterURLs(wg *sync.WaitGroup) {
	fmt.Println("Scraper started...")

	debug("Getting character URLs...", s)

	c := colly.NewCollector(colly.AllowedDomains("frieren.fandom.com"))

	c.OnHTML("div#portal_frame a[title]", func(e *colly.HTMLElement) {
		characterURL := e.Request.AbsoluteURL(e.Attr("href"))
		if _, exists := s.URLSet[characterURL]; !exists {
			s.URLSet[characterURL] = struct{}{}
			debug("Found character URL: "+characterURL, s)
			s.CharacterURLs = append(s.CharacterURLs, characterURL)
		}
	})

	c.Visit("https://frieren.fandom.com/wiki/List_of_Characters")
}

// ScrapeCharacters starts the scraping process for each character URL
func (s *Scraper) ScrapeCharacters(wg *sync.WaitGroup) {
	for _, url := range s.CharacterURLs {
		wg.Add(1)
		debug("Scraping character: "+url, s)
		go scrapeCharacter(url, wg, s.DataChannel)
	}
}

// WriteDataToCSV writes the scraped data to a CSV file
func (s *Scraper) WriteDataToCSV(filename string) error {
	// Create or open the CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	err = writer.Write([]string{"URL", "Character", "Class", "Gender", "Rank", "Species"})
	if err != nil {
		return err
	}

	// Write data for each character
	for data := range s.DataChannel {
		row := []string{
			data["url"],
			data["character"],
			data["class"],
			data["gender"],
			data["rank"],
			data["species"],
		}
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
