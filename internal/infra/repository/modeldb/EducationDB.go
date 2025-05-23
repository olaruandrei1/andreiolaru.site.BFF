package modeldb

type EducationDB struct {
	ID          uint `gorm:"primaryKey"`
	Institution string
	Degree      string
	Period      string
	Description string
	Variant     string
	PhotoPath   string
}

func (EducationDB) TableName() string {
	return "education"
}
