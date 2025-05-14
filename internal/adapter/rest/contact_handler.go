package rest

import (
	"encoding/json"
	"net/http"

	"andreiolaru.site.bff/internal/app"
	"andreiolaru.site.bff/internal/domain/model/puts"
	"andreiolaru.site.bff/internal/domain/validation"
)

type ContactHandler struct {
	Service *app.ContactService
}

func NewContactHandler(service *app.ContactService) *ContactHandler {
	return &ContactHandler{Service: service}
}

func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ContactHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	contact, err := h.Service.GetContact(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch contact", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

func (h *ContactHandler) handlePost(w http.ResponseWriter, r *http.Request) {
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
