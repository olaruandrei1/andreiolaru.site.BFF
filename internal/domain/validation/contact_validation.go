package validation

import (
	"errors"
	"regexp"
	"strings"

	puts "andreiolaru.site.bff/internal/domain/model/puts"
)

func ValidateContactMessage(msg puts.ContactMessage) error {
	if strings.TrimSpace(msg.Name) == "" {
		return errors.New("name is required")
	}
	if !isValidEmail(msg.Email) {
		return errors.New("invalid email address")
	}
	if strings.TrimSpace(msg.Message) == "" {
		return errors.New("message is required")
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return re.MatchString(email)
}
