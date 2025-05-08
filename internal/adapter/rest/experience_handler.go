package rest

import (
	"encoding/json"
	"net/http"

	"andreiolaru.site.bff/internal/app"
)

type ExperienceHandler struct {
	Service *app.ExperienceService
}

func NewExperienceHandler(service *app.ExperienceService) *ExperienceHandler {
	return &ExperienceHandler{Service: service}
}

func (h *ExperienceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetExperience(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch experience", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}
