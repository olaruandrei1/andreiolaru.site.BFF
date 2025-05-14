package modeldb

import "time"

type ContactMessageDB struct {
	ID        uint `gorm:"primaryKey"`
	Subject   string
	Email     string
	Message   string
	CreatedAt time.Time
}

func (ContactMessageDB) TableName() string {
	return "contact_messages"
}
