package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	appmodel "andreiolaru.site.bff/internal/domain/model/gets"
	dbmodel "andreiolaru.site.bff/internal/infra/repository/modeldb"
	"gorm.io/gorm"
)

type GormProjectRepository struct {
	db *gorm.DB
}

func NewGormProjectRepository(db *gorm.DB) *GormProjectRepository {
	return &GormProjectRepository{db: db}
}

func (r *GormProjectRepository) GetProjects(ctx context.Context) ([]appmodel.Project, error) {
	username := os.Getenv("GITHUB_USERNAME")
	if username == "" {
		return nil, fmt.Errorf("GITHUB_USERNAME not set")
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "GoPortfolioBackend")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var raw []struct {
		Name     string `json:"name"`
		HTMLURL  string `json:"html_url"`
		Language string `json:"language"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var projects []appmodel.Project

	for _, repo := range raw {
		lang := strings.TrimSpace(repo.Language)
		if lang == "" {
			continue
		}

		var skill dbmodel.SkillDB
		if err := r.db.WithContext(ctx).Where("skill_name = ?", lang).First(&skill).Error; err != nil {
			skill.SvgURL = "/icons/default.svg"
		}

		project := appmodel.Project{
			Title: repo.Name,
			Technologies: []appmodel.ProjectTechnology{
				{
					Name: lang,
					Icon: skill.SvgURL,
				},
			},
			RepoURL: repo.HTMLURL,
		}

		projects = append(projects, project)
	}

	return projects, nil
}
