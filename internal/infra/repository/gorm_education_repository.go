package repository

import (
	"context"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	modeldb "andreiolaru.site.bff/internal/infra/repository/modeldb"
	"gorm.io/gorm"
)

type GormEducationRepository struct {
	db *gorm.DB
}

func NewGormEducationRepository(db *gorm.DB) *GormEducationRepository {
	return &GormEducationRepository{db: db}
}

func (r *GormEducationRepository) GetEducation(ctx context.Context) ([]model.Education, error) {
	var edus []modeldb.EducationDB
	if err := r.db.WithContext(ctx).Find(&edus).Error; err != nil {
		return nil, err
	}

	var result []model.Education
	for _, e := range edus {
		result = append(result, model.Education{
			Institution: e.Institution,
			Degree:      e.Degree,
			Period:      e.Period,
			Description: e.Description,
			Variant:     model.EducationVariant(e.Variant),
			PhotoPath:   e.PhotoPath,
		})
	}

	return result, nil
}
