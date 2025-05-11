package repository

import (
	"context"

	appmodel "andreiolaru.site.bff/internal/domain/model/gets"
	dbmodel "andreiolaru.site.bff/internal/infra/repository/modeldb"
	"gorm.io/gorm"
)

type GormSkillRepository struct {
	db *gorm.DB
}

func NewGormSkillRepository(db *gorm.DB) *GormSkillRepository {
	return &GormSkillRepository{db: db}
}

func (r *GormSkillRepository) GetSkills(ctx context.Context) ([]appmodel.SkillCategory, error) {
	var dbCategories []dbmodel.SkillCategoryDB

	if err := r.db.WithContext(ctx).
		Preload("Skills", func(db *gorm.DB) *gorm.DB {
			return db.Order("order ASC")
		}).
		Order("order ASC").
		Find(&dbCategories).Error; err != nil {
		return nil, err
	}

	var result []appmodel.SkillCategory

	for _, cat := range dbCategories {
		skillsMap := make(map[string]string)
		for _, skill := range cat.Skills {
			skillsMap[skill.SkillName] = skill.SvgURL
		}

		result = append(result, appmodel.SkillCategory{
			ID:           cat.ID,
			CategoryName: cat.CategoryName,
			Skills:       skillsMap,
		})
	}

	return result, nil
}
