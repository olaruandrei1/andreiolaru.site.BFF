package seed

import (
	"context"
	"log"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedQueueSkills(db *gorm.DB) {
	ctx := context.Background()

	var existing modeldb.SkillCategoryDB
	if err := db.WithContext(ctx).Where("category_name = ?", "Queues").First(&existing).Error; err == nil {
		return
	}

	categoryID := uuid.New().String()
	category := modeldb.SkillCategoryDB{
		ID:           categoryID,
		CategoryName: "Queues",
		Order:        6,
	}

	if err := db.WithContext(ctx).Create(&category).Error; err != nil {
		log.Fatalf("failed to insert skill category: %v", err)
	}

	skills := []modeldb.SkillDB{
		{SkillName: "Kafka", SvgURL: "/svgs/kafka.svg", CategoryID: categoryID, Order: 1},
		{SkillName: "Azure ServiceBus", SvgURL: "/svgs/azure-service-bus.svg", CategoryID: categoryID, Order: 2},
		{SkillName: "AWS SQS", SvgURL: "/svgs/aws-sqs.svg", CategoryID: categoryID, Order: 3},
		{SkillName: "AWS SNS", SvgURL: "/svgs/aws-sns.svg", CategoryID: categoryID, Order: 4},
		{SkillName: "IBM-Queues", SvgURL: "/svgs/ibm-queues.svg", CategoryID: categoryID, Order: 5},
	}

	if err := db.WithContext(ctx).Create(&skills).Error; err != nil {
		log.Fatalf("failed to insert queue skills: %v", err)
	}

	log.Println("Queue skills seeded.")
}
