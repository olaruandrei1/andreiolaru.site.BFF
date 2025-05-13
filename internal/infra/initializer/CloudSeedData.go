package seed

import (
	"andreiolaru.site.bff/internal/infra/repository/modeldb"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedCloudSkills(db *gorm.DB) {
	var existing modeldb.SkillCategoryDB
	if err := db.Where("category_name = ?", "Cloud").First(&existing).Error; err == nil {
		return
	}

	categoryID := uuid.NewString()

	db.Create(&modeldb.SkillCategoryDB{
		ID:           categoryID,
		CategoryName: "Cloud",
		Order:        3,
	})

	skills := []modeldb.SkillDB{
		{SkillName: "Azure Functions", SvgURL: "/svgs/azure-functions.svg", CategoryID: categoryID, Order: 1},
		{SkillName: "Azure Key Vault", SvgURL: "/svgs/azure-key-vault.svg", CategoryID: categoryID, Order: 2},
		{SkillName: "Azure AD", SvgURL: "/svgs/azure-ad.svg", CategoryID: categoryID, Order: 3},
		{SkillName: "Azure Storage Container", SvgURL: "/svgs/azure-storage-container.svg", CategoryID: categoryID, Order: 4},
		{SkillName: "AWS Lambda", SvgURL: "/svgs/aws-lambda.svg", CategoryID: categoryID, Order: 5},
		{SkillName: "AWS S3", SvgURL: "/svgs/aws-s3.svg", CategoryID: categoryID, Order: 6},
		{SkillName: "AWS DynamoDB", SvgURL: "/svgs/aws-dynamodb.svg", CategoryID: categoryID, Order: 7},
	}

	for _, skill := range skills {
		db.Create(&skill)
	}
}
