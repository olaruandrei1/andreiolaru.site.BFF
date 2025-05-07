package app

import (
	"context"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	"andreiolaru.site.bff/internal/domain/port"
)

type MeService struct {
	repo port.MeRepository
}

func NewMeService(repo port.MeRepository) *MeService {
	return &MeService{repo: repo}
}

func (s *MeService) GetMe(ctx context.Context) (*model.Me, error) {
	return s.repo.GetMe(ctx)
}
