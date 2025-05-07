package middleware

import (
	"net/http"
	"os"
)

func ApiKeyAuthMiddleware(next http.Handler) http.Handler {
	requiredKey := os.Getenv("API_KEY")
	if requiredKey == "" {
		panic("API_KEY missing from environment")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientKey := r.Header.Get("X-BFF-Key")

		if clientKey != requiredKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
