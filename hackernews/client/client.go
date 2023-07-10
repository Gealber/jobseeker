package client

import (
	"log"
	"net/url"
	"strings"

	"github.com/Gealber/jobseeker/repositories/model"
	"github.com/gocolly/colly"
)

type client struct {
	collector *colly.Collector
}

func New(collector *colly.Collector) *client {
	return &client{collector: collector}
}

func (c *client) Search() []model.Job {
	result := make([]model.Job, 0)

	// extract job link
	// td.title span.titleline a
	// #\33 6630384 > td:nth-child(3) > span > a
	c.collector.OnHTML("tr.athing td.title span.titleline a", func(h *colly.HTMLElement) {
		link := h.Attr("href")

		// visit job link
		urlVal, _ := url.Parse(link)
		if strings.Contains(urlVal.Hostname(), YCombinatorDomain) {
			err := h.Request.Visit(link)
			if err != nil {
				log.Println(err)
			}
		} else {
			if urlVal.Hostname() != "" {
				result = append(result, model.Job{
					Link:        link,
					Description: "IN THE WEBSITE",
					Client:      ClientName,
				})
			}
		}
	})

	// extract job information
	descriptionSelector := "div.mx-auto.max-w-ycdc-page section div div.flex-grow.space-y-5 div:nth-child(4)"
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

	// Start scraping on https://news.ycombinator.com/jobs
	c.collector.Visit(urlSearch())

	return result
}

func urlSearch() string {
	return "https://news.ycombinator.com/jobs"
}

func techStack(description string) []string {
	return []string{}
}
