package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method error")
			}
			return []byte("YourSecretKey"), nil
		})

		if err != nil {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		}
	})
}
