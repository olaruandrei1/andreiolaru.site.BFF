package app

import (
	"context"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	"andreiolaru.site.bff/internal/domain/port"
)

type EducationService struct {
	repo port.EducationRepository
}

func NewEducationService(repo port.EducationRepository) *EducationService {
	return &EducationService{repo: repo}
}

func (s *EducationService) GetEducation(ctx context.Context) ([]model.Education, error) {
	return s.repo.GetEducation(ctx)
}
