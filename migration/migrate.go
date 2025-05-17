package migration

import (
	"andreiolaru.site.bff/internal/infra/repository/modeldb"
	"gorm.io/gorm"
)

func DropAllTables(db *gorm.DB) {
	db.Migrator().DropTable(
		&modeldb.AboutDB{},
		&modeldb.ContactMessageDB{},
		&modeldb.EducationDB{},
		&modeldb.ExperienceDB{},
		&modeldb.MeDB{},
		&modeldb.SkillDB{},
	)
}

func AutoMigrateAll(db *gorm.DB) {
	db.AutoMigrate(
		&modeldb.AboutDB{},
		&modeldb.ContactMessageDB{},
		&modeldb.EducationDB{},
		&modeldb.ExperienceDB{},
		&modeldb.MeDB{},
		&modeldb.SkillDB{},
	)
}
