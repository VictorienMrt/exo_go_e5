// auth.go
package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT creates a new JWT token.
// The token includes a username and an expiration time (set to 72 hours in this example).
func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "username",                            // Placeholder username, in a real application this should be dynamic.
		"exp":  time.Now().Add(time.Hour * 72).Unix(), // Token expiration time.
	})

	// The token is signed with a secret key. Replace "YourSecretKey" with a real secret key.
	tokenString, err := token.SignedString([]byte("YourSecretKey"))
	if err != nil {
		return "", err // Return an empty string and the error if the token cannot be signed.
	}

	return tokenString, nil // Return the signed token.
}

// loginHandler handles login requests.
// It generates a JWT token and sends it back to the client.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tokenString, err := GenerateJWT() // Generate a new JWT token.
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return // Return an internal server error if the token cannot be generated.
	}

	w.WriteHeader(http.StatusOK) // Set the response status to 200 OK.
	w.Write([]byte(tokenString)) // Write the token to the response body.
}
