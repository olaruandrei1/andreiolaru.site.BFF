package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	appmodel "andreiolaru.site.bff/internal/domain/model/gets"
	"andreiolaru.site.bff/internal/infra/repository/mapping"
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

	url := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=created&direction=desc", username)

	fmt.Printf("GitHub request URL: %s\n", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "GoPortfolioBackend")
	req.Header.Set("Accept", "application/vnd.github.mercy-preview+json")

	resp, err := http.DefaultClient.Do(req)
	fmt.Printf("GitHub response status: %s\n", resp.Status)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var raw []struct {
		Name     string   `json:"name"`
		HTMLURL  string   `json:"html_url"`
		Language string   `json:"language"`
		Topics   []string `json:"topics"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		fmt.Printf("JSON decode failed: %v\n", err)
		return nil, err
	}

	fmt.Printf("Decoded %d repos from GitHub\n", len(raw))

	var projects []appmodel.Project

	for _, repo := range raw {
		fmt.Printf("Processing repo: %s | Language: %s | Topics: %v\n", repo.Name, repo.Language, repo.Topics)
		var technologies []appmodel.ProjectTechnology

		lang := strings.TrimSpace(repo.Language)
		if lang != "" && !stringInSliceInsensitive(lang, repo.Topics) {
			mappedLang := mapping.NormalizeTopic(lang)
			if mappedLang != "" && !containsTech(technologies, mappedLang) {
				technologies = append(technologies, appmodel.ProjectTechnology{
					Name: mappedLang,
					Icon: r.getIconForSkill(ctx, mappedLang),
				})
			}
		}

		for _, topic := range repo.Topics {
			mapped := mapping.NormalizeTopic(topic)

			if mapped == "" {
				fmt.Printf("Topic invalid in repo %s: %q\n", repo.Name, topic)
				continue
			}

			if containsTech(technologies, mapped) {
				continue
			}

			technologies = append(technologies, appmodel.ProjectTechnology{
				Name: mapped,
				Icon: r.getIconForSkill(ctx, mapped),
			})
		}

		project := appmodel.Project{
			Title:        repo.Name,
			Technologies: technologies,
			RepoURL:      repo.HTMLURL,
		}

		projects = append(projects, project)

	}

	return projects, nil
}

func (r *GormProjectRepository) getIconForSkill(ctx context.Context, skillName string) string {
	if strings.TrimSpace(skillName) == "" {
		fmt.Println("⚠️ getIconForSkill called with empty skillName")
		return "/icons/default.svg"
	}

	fmt.Printf("Looking up icon for skill: %q\n", skillName)

	var skill dbmodel.SkillDB
	err := r.db.WithContext(ctx).
		Where("LOWER(skill_name) = ?", strings.ToLower(skillName)).
		First(&skill).Error

	if err != nil {
		fmt.Printf("Skill not found: %s (err: %v)\n", skillName, err)
		return "/icons/default.svg"
	}

	if skill.SvgURL == "" {
		fmt.Printf("Skill has no SvgURL: %s\n", skillName)
		return "/icons/default.svg"
	}

	return skill.SvgURL
}

func containsTech(techs []appmodel.ProjectTechnology, name string) bool {
	for _, tech := range techs {
		if strings.EqualFold(tech.Name, name) {
			return true
		}
	}
	return false
}

func stringInSliceInsensitive(target string, list []string) bool {
	for _, s := range list {
		if strings.EqualFold(s, target) {
			return true
		}
	}
	return false
}
