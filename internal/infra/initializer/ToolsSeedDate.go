package seed

import (
	"context"
	"log"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedToolSkills(db *gorm.DB) {
	ctx := context.Background()

	var existing modeldb.SkillCategoryDB
	if err := db.WithContext(ctx).Where("category_name = ?", "Tool & Monitoring").First(&existing).Error; err == nil {
		return
	}

	categoryID := uuid.New().String()
	category := modeldb.SkillCategoryDB{
		ID:           categoryID,
		CategoryName: "Tool & Monitoring",
		Order:        7,
	}

	if err := db.WithContext(ctx).Create(&category).Error; err != nil {
		log.Fatalf("failed to insert Tool & Monitoring category: %v", err)
	}

	skills := []modeldb.SkillDB{
		{SkillName: "Grafana", SvgURL: "/svgs/grafana.svg", CategoryID: categoryID, Order: 1},
		{SkillName: "Azure Application Insights", SvgURL: "/svgs/azure-app-insights.svg", CategoryID: categoryID, Order: 2},
		{SkillName: "Elasticsearch", SvgURL: "/svgs/elasticsearch.svg", CategoryID: categoryID, Order: 3},
		{SkillName: "KQL", SvgURL: "/svgs/kql.svg", CategoryID: categoryID, Order: 4},
	}

	if err := db.WithContext(ctx).Create(&skills).Error; err != nil {
		log.Fatalf("failed to insert Tool & Monitoring skills: %v", err)
	}

	log.Println("Tool & Monitoring skills seeded.")
}
