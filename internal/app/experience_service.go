package app

import (
	"context"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	"andreiolaru.site.bff/internal/domain/port"
)

type ExperienceService struct {
	repo port.ExperienceRepository
}

func NewExperienceService(repo port.ExperienceRepository) *ExperienceService {
	return &ExperienceService{repo: repo}
}

func (s *ExperienceService) GetExperience(ctx context.Context) ([]model.Experience, error) {
	return s.repo.GetExperience(ctx)
}
