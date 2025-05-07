package app

import (
	"context"

	model "andreiolaru.site.bff/internal/domain/model/gets"
	"andreiolaru.site.bff/internal/domain/port"
)

type ProjectService struct {
	repo port.ProjectRepository
}

func NewProjectService(repo port.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) GetProjects(ctx context.Context) ([]model.Project, error) {
	return s.repo.GetProjects(ctx)
}
