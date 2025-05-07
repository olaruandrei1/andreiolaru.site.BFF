package rest

import (
	"encoding/json"
	"net/http"

	"andreiolaru.site.bff/internal/app"
)

type MeHandler struct {
	Service *app.MeService
}

func NewMeHandler(service *app.MeService) *MeHandler {
	return &MeHandler{Service: service}
}

func (h *MeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	me, err := h.Service.GetMe(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch me", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(me)
}
