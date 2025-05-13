package seed

import (
	"context"
	"log"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedMobileSkills(db *gorm.DB) {
	ctx := context.Background()

	var existing modeldb.SkillCategoryDB
	if err := db.WithContext(ctx).Where("category_name = ?", "Mobile").First(&existing).Error; err == nil {
		return
	}

	categoryID := uuid.New().String()
	category := modeldb.SkillCategoryDB{
		ID:           categoryID,
		CategoryName: "Mobile",
		Order:        5,
	}

	if err := db.WithContext(ctx).Create(&category).Error; err != nil {
		log.Fatalf("failed to insert Mobile category: %v", err)
	}

	skills := []modeldb.SkillDB{
		{SkillName: "Swift", SvgURL: "/svgs/swift.svg", CategoryID: categoryID, Order: 1},
		{SkillName: "SwiftUI", SvgURL: "/svgs/swiftui.svg", CategoryID: categoryID, Order: 2},
		{SkillName: "XCode", SvgURL: "/svgs/xcode.svg", CategoryID: categoryID, Order: 3},
		{SkillName: "Apple", SvgURL: "/svgs/apple.svg", CategoryID: categoryID, Order: 4},
	}

	if err := db.WithContext(ctx).Create(&skills).Error; err != nil {
		log.Fatalf("failed to insert Mobile skills: %v", err)
	}

	log.Println("Mobile skills seeded.")
}
