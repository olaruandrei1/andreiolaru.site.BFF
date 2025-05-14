package seed

import (
	"context"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedOtherTopics(db *gorm.DB) error {
	var existing modeldb.SkillCategoryDB
	if err := db.Where("category_name = ?", "OtherTopics").First(&existing).Error; err == nil {
		return nil
	}
	categoryID := uuid.New().String()

	otherTopics := modeldb.SkillCategoryDB{
		ID:           categoryID,
		CategoryName: "OtherTopics",
		Order:        99,
		Skills: []modeldb.SkillDB{
			{
				SkillName:  "RABC",
				SvgURL:     "/svgs/rabc.svg",
				CategoryID: categoryID,
				Order:      1,
			},
			{
				SkillName:  "Bachelor Thesis",
				SvgURL:     "/svgs/bachelor-thesis.svg",
				CategoryID: categoryID,
				Order:      2,
			},
			{
				SkillName:  "Machine Learning",
				SvgURL:     "/svgs/machine-learning.svg",
				CategoryID: categoryID,
				Order:      3,
			},
			{
				SkillName:  "Crispy Forms",
				SvgURL:     "/svgs/crispy-forms.svg",
				CategoryID: categoryID,
				Order:      4,
			},
			{
				SkillName:  "ADO NET",
				SvgURL:     "/svgs/ado-net.svg",
				CategoryID: categoryID,
				Order:      5,
			},
			{
				SkillName:  "Discord BOT",
				SvgURL:     "/svgs/discord-bot.svg",
				CategoryID: categoryID,
				Order:      6,
			},
		},
	}

	return db.WithContext(context.Background()).Create(&otherTopics).Error
}
