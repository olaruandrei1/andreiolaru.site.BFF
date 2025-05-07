package app

import (
	"context"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	"andreiolaru.site.bff/internal/domain/port"
)

type SkillService struct {
	repo port.SkillRepository
}

func NewSkillService(repo port.SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) GetSkills(ctx context.Context) ([]model.SkillCategory, error) {
	return s.repo.GetSkills(ctx)
}
