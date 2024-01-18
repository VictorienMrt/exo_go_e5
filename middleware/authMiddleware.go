package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware is a middleware function for JWT authentication.
// It checks for a valid JWT token in the Authorization header of the request.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the Authorization header.
		authHeader := r.Header.Get("Authorization")
		// Split the header to separate the 'Bearer' part from the token itself.
		bearerToken := strings.Split(authHeader, " ")

		// Check if the token is correctly formatted.
		if len(bearerToken) != 2 {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		// Parse the token.
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			// Ensure the token uses the expected signing method (HMAC in this case).
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method error")
			}
			// Return the signing key.
			return []byte("YourSecretKey"), nil
		})

		// Handle any error that occurred in parsing the token.
		if err != nil {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid.
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// If valid, proceed to the next handler.
			next.ServeHTTP(w, r)
		} else {
			// If not valid, return an unauthorized error.
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		}
	})
}
