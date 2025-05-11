package repository

import (
	"context"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	modeldb "andreiolaru.site.bff/internal/infra/repository/modeldb"
	"gorm.io/gorm"
)

type GormExperienceRepository struct {
	db *gorm.DB
}

func NewGormExperienceRepository(db *gorm.DB) *GormExperienceRepository {
	return &GormExperienceRepository{db: db}
}

func (r *GormExperienceRepository) GetExperience(ctx context.Context) ([]model.Experience, error) {
	var dbExps []modeldb.ExperienceDB
	if err := r.db.WithContext(ctx).Find(&dbExps).Error; err != nil {
		return nil, err
	}

	var exps []model.Experience
	for _, dbExp := range dbExps {
		exps = append(exps, model.Experience{
			Company:     dbExp.Company,
			Title:       dbExp.Title,
			Period:      dbExp.Period,
			Description: dbExp.Description,
			PhotoPath:   dbExp.PhotoPath,
		})
	}

	return exps, nil
}
