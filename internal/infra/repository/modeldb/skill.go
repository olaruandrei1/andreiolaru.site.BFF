package modeldb

type SkillCategoryDB struct {
	ID           string    `gorm:"primaryKey;column:id;size:36"`
	CategoryName string    `gorm:"column:category_name"`
	Order        uint      `gorm:"column:order"`
	Skills       []SkillDB `gorm:"foreignKey:CategoryID"`
}

func (SkillCategoryDB) TableName() string {
	return "skill_category"
}

type SkillDB struct {
	ID         uint   `gorm:"primaryKey"`
	SkillName  string `gorm:"column:skill_name"`
	SvgURL     string `gorm:"column:svg_url;type:LONGTEXT"`
	CategoryID string `gorm:"column:category_id;size:36"`
	Order      uint   `gorm:"column:order"`
}

func (SkillDB) TableName() string {
	return "skill"
}
