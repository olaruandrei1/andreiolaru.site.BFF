package rest

import (
	"encoding/json"
	"net/http"

	"andreiolaru.site.bff/internal/app"
	"andreiolaru.site.bff/internal/domain/model/puts"
	"andreiolaru.site.bff/internal/domain/validation"
)

type ContactPostHandler struct {
	Service *app.ContactService
}

func NewContactPostHandler(service *app.ContactService) *ContactPostHandler {
	return &ContactPostHandler{Service: service}
}

func (h *ContactPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var payload puts.ContactMessage

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := validation.ValidateContactMessage(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.SaveMessage(r.Context(), payload); err != nil {
		http.Error(w, "failed to save message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Message received"})
}
