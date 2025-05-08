package migration

import (
	"log"

	"gorm.io/gorm"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"
)

func AutoMigrateAll(db *gorm.DB) {
	err := db.AutoMigrate(
		&modeldb.MeDB{},
		&modeldb.AboutDB{},
		&modeldb.ExperienceDB{},
		&modeldb.SkillCategoryDB{},
		&modeldb.SkillDB{},
		&modeldb.EducationDB{},
		&modeldb.ContactMessageDB{},
	)

	if err != nil {
		log.Fatalf("Auto-migrate failed: %v", err)
	}

	log.Println("All models migrated successfully.")
}

func DropAllTables(db *gorm.DB) {
	err := db.Migrator().DropTable(
		&modeldb.MeDB{},
		&modeldb.AboutDB{},
		&modeldb.ExperienceDB{},
		&modeldb.SkillCategoryDB{},
		&modeldb.SkillDB{},
		&modeldb.EducationDB{},
		&modeldb.ContactMessageDB{},
	)

	if err != nil {
		log.Printf("⚠️  Failed to drop tables: %v", err)
	} else {
		log.Println("🧹 All tables dropped successfully.")
	}
}
