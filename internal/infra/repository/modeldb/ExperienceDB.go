package modeldb

type ExperienceDB struct {
	ID          uint `gorm:"primaryKey"`
	Company     string
	Title       string
	Period      string
	Description string
}

func (ExperienceDB) TableName() string {
	return "experience"
}
