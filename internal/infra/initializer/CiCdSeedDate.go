package seed

import (
	"context"
	"log"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedCICDSkills(db *gorm.DB) {
	ctx := context.Background()

	var existing modeldb.SkillCategoryDB
	if err := db.WithContext(ctx).Where("category_name = ?", "CI/CD & Methodologies").First(&existing).Error; err == nil {
		return
	}

	categoryID := uuid.New().String()
	category := modeldb.SkillCategoryDB{
		ID:           categoryID,
		CategoryName: "CI/CD & Methodologies",
	}

	if err := db.WithContext(ctx).Create(&category).Error; err != nil {
		log.Fatalf("failed to insert CI/CD category: %v", err)
	}

	skills := []modeldb.SkillDB{
		{SkillName: "Git", SvgURL: "/svgs/git.svg", CategoryID: categoryID},
		{SkillName: "GitHub", SvgURL: "/svgs/github.svg", CategoryID: categoryID},
		{SkillName: "Azure DevOps", SvgURL: "/svgs/azure-devops.svg", CategoryID: categoryID},
		{SkillName: "Jenkins", SvgURL: "/svgs/jenkins.svg", CategoryID: categoryID},
		{SkillName: "Jira", SvgURL: "/svgs/jira.svg", CategoryID: categoryID},
		{SkillName: "Confluence", SvgURL: "/svgs/confluence.svg", CategoryID: categoryID},
		{SkillName: "Agile", SvgURL: "/svgs/agile.svg", CategoryID: categoryID},
		{SkillName: "Scrum", SvgURL: "/svgs/scrum.svg", CategoryID: categoryID},
		{SkillName: "Kanban", SvgURL: "/svgs/kanban.svg", CategoryID: categoryID},
		{SkillName: "Waterfall", SvgURL: "/svgs/waterfall.svg", CategoryID: categoryID},
	}

	if err := db.WithContext(ctx).Create(&skills).Error; err != nil {
		log.Fatalf("failed to insert CI/CD skills: %v", err)
	}

	log.Println("CI/CD & Methodologies skills seeded.")
}
