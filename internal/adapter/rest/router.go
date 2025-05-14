package rest

import (
	"net/http"

	"andreiolaru.site.bff/internal/app"
)

type Router struct {
	MeHandler         *MeHandler
	AboutHandler      *AboutHandler
	ExperienceHandler *ExperienceHandler
	SkillHandler      *SkillHandler
	ProjectHandler    *ProjectHandler
	EducationHandler  *EducationHandler
	ContactHandler    *ContactHandler
}

func NewRouter(
	meService *app.MeService,
	aboutService *app.AboutService,
	experienceService *app.ExperienceService,
	skillService *app.SkillService,
	projectService *app.ProjectService,
	educationService *app.EducationService,
	contactService *app.ContactService,
) *Router {
	return &Router{
		MeHandler:         NewMeHandler(meService),
		AboutHandler:      NewAboutHandler(aboutService),
		ExperienceHandler: NewExperienceHandler(experienceService),
		SkillHandler:      NewSkillHandler(skillService),
		ProjectHandler:    NewProjectHandler(projectService),
		EducationHandler:  NewEducationHandler(educationService),
		ContactHandler:    NewContactHandler(contactService),
	}
}

func (r *Router) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/me", r.MeHandler)
	mux.Handle("/about", r.AboutHandler)
	mux.Handle("/experience", r.ExperienceHandler)
	mux.Handle("/skills", r.SkillHandler)
	mux.Handle("/projects", r.ProjectHandler)
	mux.Handle("/education", r.EducationHandler)
	mux.Handle("/contact", r.ContactHandler)

	return mux
}
