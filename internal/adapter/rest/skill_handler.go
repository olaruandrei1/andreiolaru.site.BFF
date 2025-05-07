package rest

import (
	"encoding/json"
	"net/http"

	"andreiolaru.site.bff/internal/app"
)

type SkillHandler struct {
	Service *app.SkillService
}

func NewSkillHandler(service *app.SkillService) *SkillHandler {
	return &SkillHandler{Service: service}
}

func (h *SkillHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetSkills(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch skills", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}
