package modeldb

type AboutDB struct {
	ID           uint `gorm:"primaryKey"`
	WhoIAmTitle  string
	WhoIAmText   string
	MindsetTitle string
	MindsetText  string
	CVURL        string
}

func (AboutDB) TableName() string {
	return "about"
}
