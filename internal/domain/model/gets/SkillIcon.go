package model

type SkillIcon struct {
	ID        uint   `gorm:"primaryKey"`
	SkillName string `gorm:"column:skill_name"`
	SvgURL    string `gorm:"column:svg_url"`
}
