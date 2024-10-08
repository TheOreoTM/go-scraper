# Go Scraper
A simple web scraper written in Go to scrape data from the Frieren Fandom website.

## Usage

```bash
go run main.go
```

## Features
- Scrapes the Frieren Fandom website for all entries
- Automatically crawls all pages of characters
- Collects important data from each entry
    - [x] Name
    - [x] Description
    - [x] Abilities
    - [ ] Skills
    - [x] Class
    - [x] Gender
    - [x] Rank
    - [x] Species
- Parses anchor tags and converts them to hyperlinks to be used in discord
- Outputs the data in JSON

## Output

![Output showing the scraped data](assets/data_preview.gif)
