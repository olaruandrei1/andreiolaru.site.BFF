package rest

import (
	"encoding/json"
	"net/http"

	"andreiolaru.site.bff/internal/app"
)

type AboutHandler struct {
	Service *app.AboutService
}

func NewAboutHandler(service *app.AboutService) *AboutHandler {
	return &AboutHandler{Service: service}
}

func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	about, err := h.Service.GetAbout(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch about", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(about)
}
