package seed

import (
	"context"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedDatabaseSkills(db *gorm.DB) error {
	var existing modeldb.SkillCategoryDB
	if err := db.Where("category_name = ?", "Database").First(&existing).Error; err == nil {
		return nil
	}

	categoryID := uuid.New().String()

	databaseCategory := modeldb.SkillCategoryDB{
		ID:           categoryID,
		CategoryName: "Database",
		Skills: []modeldb.SkillDB{
			{
				SkillName:  "Microsoft SQL Server (T-SQL)",
				SvgURL:     "/svgs/microsoft-sql-server.svg",
				CategoryID: categoryID,
			},
			{
				SkillName:  "Oracle (PL/SQL)",
				SvgURL:     "/svgs/oracle.svg",
				CategoryID: categoryID,
			},
			{
				SkillName:  "PostgreSQL (PL/pgSQL)",
				SvgURL:     "/svgs/postgresql.svg",
				CategoryID: categoryID,
			},
			{
				SkillName:  "MySQL",
				SvgURL:     "/svgs/mysql.svg",
				CategoryID: categoryID,
			},
			{
				SkillName:  "SQLite",
				SvgURL:     "/svgs/sqlite.svg",
				CategoryID: categoryID,
			},
		},
	}

	return db.WithContext(context.Background()).Create(&databaseCategory).Error
}
