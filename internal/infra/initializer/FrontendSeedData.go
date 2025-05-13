package seed

import (
	"context"
	"log"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedFrontendSkills(db *gorm.DB) {

	ctx := context.Background()

	var existing modeldb.SkillCategoryDB
	if err := db.WithContext(ctx).Where("category_name = ?", "Frontend").First(&existing).Error; err == nil {
		return
	}

	categoryID := uuid.New().String()
	category := modeldb.SkillCategoryDB{
		ID:           categoryID,
		CategoryName: "Frontend",
		Order:        2,
	}

	if err := db.WithContext(ctx).Create(&category).Error; err != nil {
		log.Fatalf("failed to insert skill category: %v", err)
	}

	skills := []modeldb.SkillDB{
		{SkillName: "React", SvgURL: "/svgs/react.svg", CategoryID: categoryID, Order: 1},
		{SkillName: "Angular", SvgURL: "/svgs/angular.svg", CategoryID: categoryID, Order: 2},
		{SkillName: "Vue", SvgURL: "/svgs/vue.svg", CategoryID: categoryID, Order: 3},
		{SkillName: "Blazor", SvgURL: "/svgs/blazor.svg", CategoryID: categoryID, Order: 4},
		{SkillName: "Vite", SvgURL: "/svgs/vite.svg", CategoryID: categoryID, Order: 5},
		{SkillName: "JavaScript", SvgURL: "/svgs/javascript.svg", CategoryID: categoryID, Order: 6},
		{SkillName: "TypeScript", SvgURL: "/svgs/typescript.svg", CategoryID: categoryID, Order: 7},
		{SkillName: "HTML", SvgURL: "/svgs/html.svg", CategoryID: categoryID, Order: 8},
		{SkillName: "CSS", SvgURL: "/svgs/css.svg", CategoryID: categoryID, Order: 9},
		{SkillName: "Tailwind", SvgURL: "/svgs/tailwind.svg", CategoryID: categoryID, Order: 10},
		{SkillName: "Material UI", SvgURL: "/svgs/material-ui.svg", CategoryID: categoryID, Order: 11},
		{SkillName: "jQuery", SvgURL: "/svgs/jquery.svg", CategoryID: categoryID, Order: 12},
	}

	if err := db.WithContext(ctx).Create(&skills).Error; err != nil {
		log.Fatalf("failed to insert skills: %v", err)
	}

	log.Println("Frontend skills seeded.")
}
