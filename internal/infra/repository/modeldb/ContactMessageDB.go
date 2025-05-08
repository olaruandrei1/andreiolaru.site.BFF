package modeldb

import "time"

type ContactMessageDB struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Message   string
	CreatedAt time.Time
}

func (ContactMessageDB) TableName() string {
	return "contact_messages"
}
