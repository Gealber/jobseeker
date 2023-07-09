package job

import (
	"context"

	"github.com/Gealber/jobseeker/repositories/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, job model.Job) (*model.Job, error) {
	// in case we try to insert a job that is already in db just update the value.
	if err := r.db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "link"}},
			DoNothing: true,
		}).
		Create(&job).Error; err != nil {
		return nil, err
	}

	return &job, nil
}

func (r *repository) Bulk(ctx context.Context, jobs []model.Job) ([]model.Job, error) {
	// in case we try to insert a job that is already in db just update the value.
	if err := r.db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "link"}},
			DoNothing: true,
		}).
		Create(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *repository) Find(ctx context.Context) ([]model.Job, error) {
	var jobs []model.Job
	err := r.db.
		Find(&jobs).Error
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	return r.db.
		Where("id = ?", id).
		Delete(&model.Job{}).Error
}
