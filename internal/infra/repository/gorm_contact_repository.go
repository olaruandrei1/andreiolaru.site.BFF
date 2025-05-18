package repository

import (
	"context"
	"time"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	puts "andreiolaru.site.bff/internal/domain/model/puts"
	modeldb "andreiolaru.site.bff/internal/infra/repository/modeldb"
	"gorm.io/gorm"
)

type GormContactRepository struct {
	db *gorm.DB
}

func NewGormContactRepository(db *gorm.DB) *GormContactRepository {
	return &GormContactRepository{db: db}
}

func (r *GormContactRepository) GetContact(ctx context.Context) (*model.Contact, error) {
	return &model.Contact{
		LinkedIn: "https://www.linkedin.com/in/andrei-olaru-23504b362/",
		GitHub:   "https://github.com/olaruandrei1",
		X:        "https://x.com/AndreiOlar78690",
	}, nil
}

func (r *GormContactRepository) SaveContactMessage(ctx context.Context, msg puts.ContactMessage) error {
	return r.db.WithContext(ctx).Create(&modeldb.ContactMessageDB{
		Subject:   msg.Subject,
		Email:     msg.Email,
		Message:   msg.Message,
		CreatedAt: time.Now(),
	}).Error
}
