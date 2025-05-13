package repository

import (
	"context"
	"log"

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
			return db.Order("`skill`.`order` ASC")
		}).
		Order("`skill_category`.`order` ASC").
		Find(&dbCategories).Error; err != nil {
		log.Printf("Error fetching skills: %v", err)
		return nil, err
	}

	var result []appmodel.SkillCategory

	for _, cat := range dbCategories {
		var skills []appmodel.Skill
		for _, skill := range cat.Skills {
			skills = append(skills, appmodel.Skill{
				Name:   skill.SkillName,
				SvgURL: skill.SvgURL,
				Order:  skill.Order,
			})
		}

		result = append(result, appmodel.SkillCategory{
			ID:           cat.ID,
			CategoryName: cat.CategoryName,
			Order:        cat.Order,
			Skills:       skills,
		})
	}

	return result, nil
}
