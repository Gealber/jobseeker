package client

import (
	"fmt"
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

func (c *client) Search(params []SearchParam) []model.Job {
	result := make([]model.Job, 0)

	// extract job link
	c.collector.OnHTML("div.yuRUbf a", func(h *colly.HTMLElement) {
		link := h.Attr("href")

		// cleaning link
		if strings.Contains(link, "translate.google.com") {
			q, err := url.ParseQuery(link)
			if err != nil {
				// ignoring this link
				return
			}

			link = q.Get("u")
		}

		result = append(result, model.Job{
			Link: link,
		})
	})

	// Before making a request print "Visiting ..."
	c.collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	for _, param := range params {
		if !param.Valid() {
			log.Println("ignoring param, please check Period in param...")
			continue
		}

		c.collector.Visit(urlSearch(param))
	}

	return result
}

func urlSearch(s SearchParam) string {
	return fmt.Sprintf("https://www.google.com/search?%s", s.Query())
}

func techStack(description string) []string {
	return []string{}
}
