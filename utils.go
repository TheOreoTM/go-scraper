package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Helper function to clean up text, removing unwanted tags like <br> or <sup>
func cleanText(selection *goquery.Selection) string {
	var rankText strings.Builder

	selection.Contents().Each(func(i int, s *goquery.Selection) {
		if s.Is("sup") {
			return // Ignore <sup> elements
		}
		if s.Is("br") {
			rankText.WriteString("\n") // Handle <br> as newlines
			return
		}
		if s.Is("a") {
			href, exists := s.Attr("href")
			if exists {
				// Format hyperlink as Markdown: [text](link)
				fmt.Println(s.Text())
				rankText.WriteString(fmt.Sprintf("[%s](%s)", strings.TrimSpace(s.Text()), fmt.Sprintf("https://frieren.fandom.com%s", href)))
			}
			return
		}
	})

	// Finally, remove any trailing spaces
	return strings.TrimSpace(rankText.String())
}

func debug(msg string, s *Scraper) {
	if s.ShouldDebug {
		fmt.Println(msg)
	}
}
