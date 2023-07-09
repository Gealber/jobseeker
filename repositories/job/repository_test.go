package job_test

import (
	"context"
	"testing"

	"github.com/Gealber/jobseeker/config"
	jobRepo "github.com/Gealber/jobseeker/repositories/job"
	"github.com/Gealber/jobseeker/repositories/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_Create(t *testing.T) {
	// open connection with db in the test database
	ctx := context.Background()
	cfg := config.Config()

	db, err := gorm.Open(postgres.Open(cfg.Database.TESTDSN), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to connect database", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			t.Fatal(err)
		}

		sqlDB.Close()
	}()

	repo := jobRepo.NewRepository(db)

	t.Run("create job", func(t *testing.T) {
		job := model.Job{
			Link:        "https://example.com/job/1",
			Description: "Golang Developer required",
			TechStack:   []string{"Golang"},
		}

		jobCreated, err := repo.Create(ctx, job)
		if err != nil {
			t.Fatal(err)
		}

		if jobCreated.ID == "" {
			t.Fatal("job id is zero")
		}

		db.Delete(jobCreated)
	})
}

func Test_Bulk(t *testing.T) {
	// open connection with db in the test database
	ctx := context.Background()
	cfg := config.Config()

	db, err := gorm.Open(postgres.Open(cfg.Database.TESTDSN), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to connect database", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			t.Fatal(err)
		}

		sqlDB.Close()
	}()

	repo := jobRepo.NewRepository(db)

	t.Run("bulk create job", func(t *testing.T) {
		jobs := make([]model.Job, 0)
		jobs = append(jobs, model.Job{
			Link:        "https://example.com/job/1",
			Description: "Golang Developer required",
			TechStack:   []string{"Golang"},
		})

		jobsCreated, err := repo.Bulk(ctx, jobs)
		if err != nil {
			t.Fatal(err)
		}

		for _, jobCreated := range jobsCreated {
			if jobCreated.ID == "" {
				t.Fatal("job id is zero")
			}
		}

		db.Delete(jobsCreated)
	})
}
