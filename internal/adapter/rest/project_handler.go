package rest

import (
	"encoding/json"
	"net/http"

	"andreiolaru.site.bff/internal/app"
)

type ProjectHandler struct {
	Service *app.ProjectService
}

func NewProjectHandler(service *app.ProjectService) *ProjectHandler {
	return &ProjectHandler{Service: service}
}

func (h *ProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Service.GetProjects(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch projects", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(projects)
}
