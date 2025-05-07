package rest

import (
	"encoding/json"
	"net/http"

	"andreiolaru.site.bff/internal/app"
)

type ContactHandler struct {
	Service *app.ContactService
}

func NewContactHandler(service *app.ContactService) *ContactHandler {
	return &ContactHandler{Service: service}
}

func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contact, err := h.Service.GetContact(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch contact", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(contact)
}
