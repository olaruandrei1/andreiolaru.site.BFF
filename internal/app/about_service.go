package app

import (
	"context"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	"andreiolaru.site.bff/internal/domain/port"
)

type AboutService struct {
	repo port.AboutRepository
}

func NewAboutService(repo port.AboutRepository) *AboutService {
	return &AboutService{repo: repo}
}

func (s *AboutService) GetAbout(ctx context.Context) (*model.About, error) {
	return s.repo.GetAbout(ctx)
}
