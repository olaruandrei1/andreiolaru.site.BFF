package repository

import (
	"context"
	"errors"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	modeldb "andreiolaru.site.bff/internal/infra/repository/modeldb"
	"gorm.io/gorm"
)

type GormMeRepository struct {
	db *gorm.DB
}

func NewGormMeRepository(db *gorm.DB) *GormMeRepository {
	return &GormMeRepository{db: db}
}

func (r *GormMeRepository) GetMe(ctx context.Context) (*model.Me, error) {
	var dbMe modeldb.MeDB
	if err := r.db.WithContext(ctx).First(&dbMe).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &model.Me{
		Title:       dbMe.Title,
		Job:         dbMe.Job,
		Description: dbMe.Description,
		ImageURL:    dbMe.ImageURL,
	}, nil
}
