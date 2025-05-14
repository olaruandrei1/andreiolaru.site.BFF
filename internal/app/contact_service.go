package app

import (
	"context"
	"fmt"
	"strings"

	"andreiolaru.site.bff/internal/app/helpers/httpclient"
	gets "andreiolaru.site.bff/internal/domain/model/gets"
	puts "andreiolaru.site.bff/internal/domain/model/puts"
	"andreiolaru.site.bff/internal/domain/port"
)

type ContactService struct {
	repo       port.ContactRepository
	httpClient *httpclient.HTTPClientService
	apiURL     string
	apiKey     string
}

func NewContactService(repo port.ContactRepository, client *httpclient.HTTPClientService, apiURL string, apiKey string) *ContactService {
	return &ContactService{
		repo:       repo,
		httpClient: client,
		apiURL:     apiURL,
		apiKey:     apiKey,
	}
}

func (s *ContactService) GetContact(ctx context.Context) (*gets.Contact, error) {
	return s.repo.GetContact(ctx)
}

func (s *ContactService) SaveMessage(ctx context.Context, msg puts.ContactMessage) error {
	trimmed := puts.ContactMessage{
		Subject: strings.TrimSpace(msg.Subject),
		Email:   strings.TrimSpace(msg.Email),
		Message: strings.TrimSpace(msg.Message),
	}

	if err := s.repo.SaveContactMessage(ctx, trimmed); err != nil {
		return fmt.Errorf("failed to save contact message locally: %w", err)
	}

	payload := map[string]string{
		"From":    trimmed.Email,
		"Subject": trimmed.Subject,
		"Message": trimmed.Message,
	}

	headers := map[string]string{
		"X-API-Key": s.apiKey,
	}

	resp, err := s.httpClient.PostJSON(ctx, s.apiURL, payload, headers)

	if err != nil {
		return fmt.Errorf("failed to send message to external API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("external API returned status: %s", resp.Status)
	}

	return nil
}
