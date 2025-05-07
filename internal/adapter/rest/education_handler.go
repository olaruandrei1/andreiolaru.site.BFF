package rest

import (
	"encoding/json"
	"net/http"

	"andreiolaru.site.bff/internal/app"
)

type EducationHandler struct {
	Service *app.EducationService
}

func NewEducationHandler(service *app.EducationService) *EducationHandler {
	return &EducationHandler{Service: service}
}

func (h *EducationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	education, err := h.Service.GetEducation(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch education", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(education)
}
