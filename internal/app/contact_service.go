package app

import (
	"context"

	gets "andreiolaru.site.bff/internal/domain/model/gets"
	puts "andreiolaru.site.bff/internal/domain/model/puts"
	"andreiolaru.site.bff/internal/domain/port"
)

type ContactService struct {
	repo port.ContactRepository
}

func NewContactService(repo port.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) GetContact(ctx context.Context) (*gets.Contact, error) {
	return s.repo.GetContact(ctx)
}

func (s *ContactService) SaveMessage(ctx context.Context, msg puts.ContactMessage) error {
	return s.repo.SaveContactMessage(ctx, msg)
}
