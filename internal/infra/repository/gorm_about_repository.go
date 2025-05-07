package repository

import (
	"context"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	modeldb "andreiolaru.site.bff/internal/infra/repository/modeldb"
	"gorm.io/gorm"
)

type GormAboutRepository struct {
	db *gorm.DB
}

func NewGormAboutRepository(db *gorm.DB) *GormAboutRepository {
	return &GormAboutRepository{db: db}
}

func (r *GormAboutRepository) GetAbout(ctx context.Context) (*model.About, error) {
	var about modeldb.AboutDB
	if err := r.db.WithContext(ctx).First(&about).Error; err != nil {
		return nil, err
	}

	return &model.About{
		WhoIAm: model.AboutBlock{
			Title:   about.WhoIAmTitle,
			Content: about.WhoIAmText,
		},
		Mindset: model.AboutBlock{
			Title:   about.MindsetTitle,
			Content: about.MindsetText,
		},
		CVURL: about.CVURL,
	}, nil
}
