package seed

import (
	"context"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedBackendSkills(db *gorm.DB) error {
	var existing modeldb.SkillCategoryDB
	if err := db.Where("category_name = ?", "Backend").First(&existing).Error; err == nil {
		return nil
	}
	categoryID := uuid.New().String()

	backendCategory := modeldb.SkillCategoryDB{
		ID:           categoryID,
		CategoryName: "Backend",
		Order:        1,
		Skills: []modeldb.SkillDB{
			{SkillName: "C#", SvgURL: "/svgs/csharp.svg", CategoryID: categoryID, Order: 1},
			{SkillName: ".NET", SvgURL: "/svgs/dotnet.svg", CategoryID: categoryID, Order: 2},
			{SkillName: "MediatR", SvgURL: "/svgs/mediatr.svg", CategoryID: categoryID, Order: 3},
			{SkillName: "Entity Framework", SvgURL: "/svgs/ef.svg", CategoryID: categoryID, Order: 4},
			{SkillName: "API", SvgURL: "/svgs/api.svg", CategoryID: categoryID, Order: 5},
			{SkillName: "Go", SvgURL: "/svgs/go.svg", CategoryID: categoryID, Order: 6},
			{SkillName: "Gorm-ORM", SvgURL: "/svgs/gorm.svg", CategoryID: categoryID, Order: 7},
			{SkillName: "Java", SvgURL: "/svgs/java.svg", CategoryID: categoryID, Order: 8},
			{SkillName: "Spring Boot", SvgURL: "/svgs/spring-boot.svg", CategoryID: categoryID, Order: 9},
			{SkillName: "Python", SvgURL: "/svgs/python.svg", CategoryID: categoryID, Order: 10},
			{SkillName: "Django", SvgURL: "/svgs/django.svg", CategoryID: categoryID, Order: 11},
			{SkillName: "C++", SvgURL: "/svgs/cpp.svg", CategoryID: categoryID, Order: 12},
			{SkillName: "STL", SvgURL: "/svgs/standard-template-library.svg", CategoryID: categoryID, Order: 13},
		},
	}

	return db.WithContext(context.Background()).Create(&backendCategory).Error
}
