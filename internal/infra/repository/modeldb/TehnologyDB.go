package modeldb

type TechnologyDB struct {
	ID        uint   `gorm:"primaryKey"`
	SkillName string `gorm:"uniqueIndex"`
	SvgURL    string
}
