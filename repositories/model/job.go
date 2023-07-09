package model

import (
	"time"

	"github.com/lib/pq"
)

type Job struct {
	ID          string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Link        string         `gorm:"column:link"`
	Description string         `gorm:"column:description"`
	TechStack   pq.StringArray `gorm:"column:tech_stack;type:varchar(255)[]"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
}

func (Job) TableName() string {
	return "jobs"
}
