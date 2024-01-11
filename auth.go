// auth.go
package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "username",
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("YourSecretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	tokenString, err := GenerateJWT()
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}
