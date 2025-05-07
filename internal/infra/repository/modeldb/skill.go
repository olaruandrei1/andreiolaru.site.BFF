package modeldb

type SkillCategoryDB struct {
	ID           string    `gorm:"primaryKey;column:id"`
	CategoryName string    `gorm:"column:category_name"`
	Skills       []SkillDB `gorm:"foreignKey:CategoryID"`
}

type SkillDB struct {
	ID         uint   `gorm:"primaryKey"`
	SkillName  string `gorm:"column:skill_name"`
	SvgURL     string `gorm:"column:svg_url"`
	CategoryID string `gorm:"column:category_id"`
}
