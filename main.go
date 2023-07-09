package main

import (
	"context"
	"log"

	"github.com/Gealber/jobseeker/config"
	// linkedinClient "github.com/Gealber/jobseeker/linkedin/client"
	ycombinatorClient "github.com/Gealber/jobseeker/hackernews/client"
	jobRepo "github.com/Gealber/jobseeker/repositories/job"

	"github.com/gocolly/colly"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// open connection with db in the test database
	ctx := context.Background()
	cfg := config.Config()

	db, err := gorm.Open(postgres.Open(cfg.Database.TESTDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}

		sqlDB.Close()
	}()

	// Instantiate default collector
	collector := colly.NewCollector(
		// Visit only domains: linkedin.com, and www.linkedin.com
		colly.AllowedDomains(
			"linkedin.com",
			"www.linkedin.com",
			"news.ycombinator.com",
			"ycombinator.com",
			"www.ycombinator.com",
		),
		// User agent
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"),
	)

	// linkedinClt := linkedinClient.New(collector)
	// params := []linkedinClient.SearchParam{
	// 	{Keywords: "Golang", Location: "United States", FTPR: "r604800", FWT: "2", Position: 1, PageNum: 0},
	// }
	// jobs := linkedinClt.Search(params)

	ycombinatorClt := ycombinatorClient.New(collector)
	jobs := ycombinatorClt.Search()
	repo := jobRepo.NewRepository(db)

	if len(jobs) > 0 {
		if _, err := repo.Bulk(ctx, jobs); err != nil {
			log.Fatal(err)
		}
	}
}
