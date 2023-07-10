package client

import (
	"fmt"
	"log"

	"github.com/Gealber/jobseeker/repositories/model"
	"github.com/gocolly/colly"
)

type client struct {
	collector *colly.Collector
}

func New(collector *colly.Collector) *client {
	return &client{collector: collector}
}

func (c *client) Search(params []SearchParam) []model.Job {
	result := make([]model.Job, 0)

	// extract job link
	c.collector.OnHTML("ul.jobs-search__results-list div.base-card a.base-card__full-link", func(h *colly.HTMLElement) {
		link := h.Attr("href")

		// visit job link
		h.Request.Visit(link)
	})

	// extract job information
	descriptionSelector := "#main-content section div div section div div section.show-more-less-html div"
	c.collector.OnHTML(descriptionSelector, func(h *colly.HTMLElement) {
		result = append(result, model.Job{
			Link:        h.Request.URL.String(),
			Description: h.Text,
			Client:      ClientName,
		})
	})

	// Before making a request print "Visiting ..."
	c.collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	for _, param := range params {
		if !param.Valid() {
			log.Println("ignoring param, please check FWT in param...")
			continue
		}

		// Start scraping on https://linkedin.com
		c.collector.Visit(urlSearch(param))
	}

	return result
}

func urlSearch(param SearchParam) string {
	return fmt.Sprintf("%s%s?%s", Domain, DefaultPath, param.Query())
}

func techStack(description string) []string {
	return []string{}
}
