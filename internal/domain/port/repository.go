package port

import (
	"context"

	gets "andreiolaru.site.bff/internal/domain/model/gets"
	puts "andreiolaru.site.bff/internal/domain/model/puts"
)

type MeRepository interface {
	GetMe(ctx context.Context) (*gets.Me, error)
}

type AboutRepository interface {
	GetAbout(ctx context.Context) (*gets.About, error)
}

type ExperienceRepository interface {
	GetExperience(ctx context.Context) ([]gets.Experience, error)
}

type SkillRepository interface {
	GetSkills(ctx context.Context) ([]gets.SkillCategory, error)
}

type ProjectRepository interface {
	GetProjects(ctx context.Context) ([]gets.Project, error)
}

type EducationRepository interface {
	GetEducation(ctx context.Context) ([]gets.Education, error)
}

type ContactRepository interface {
	GetContact(ctx context.Context) (*gets.Contact, error)
	SaveContactMessage(ctx context.Context, msg puts.ContactMessage) error
}
