package main

import (
	"log"

	"github.com/Gealber/jobseeker/linkedin/client"
	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	collector := colly.NewCollector(
		// Visit only domains: linkedin.com, and www.linkedin.com
		colly.AllowedDomains("linkedin.com", "www.linkedin.com"),
		// User agent
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"),
	)

	clt := client.New(collector)
	params := []client.SearchParam{
		{Keywords: "Golang", Location: "United States", FTPR: "r604800", FWT: "2", Position: 1, PageNum: 0},
	}

	jobs := clt.Search(params)

	for _, job := range jobs {
		log.Println("FOUND: ", job)
	}
}
