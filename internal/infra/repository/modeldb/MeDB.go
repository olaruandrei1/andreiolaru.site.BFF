package modeldb

type MeDB struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Job         string
	Description string
	ImageURL    string
}

func (MeDB) TableName() string {
	return "me"
}
