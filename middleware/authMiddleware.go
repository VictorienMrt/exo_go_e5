package middleware

import (
	"net/http"
)

const authToken = "secret_token"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != authToken {
			http.Error(w, "Acces denied", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
