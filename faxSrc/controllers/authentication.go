package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"thermalFax/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func isAuthenticated(r *http.Request) bool {
	cookies := r.Cookies()
	for _, c := range cookies {
		if c.Name == "token" {
			// Check if the token exists and if its date is valid
			expirationTime := time.Now()
			return expirationTime.Unix() >= c.Expires.Unix()
		}
	}

	return false
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	// first we have to verify if this request brings any login data
	// Otherwise we login, which we will do everytime lmao

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	var loginForm models.LoginForm
	err = json.Unmarshal(data, &loginForm)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	if !models.ValidateCredentials(&loginForm) {
		log.Printf("Invalid credentials.")
		return
	}

	// Generate the thing
	expirationTime := time.Now().Add(2 * time.Minute)
	claims := models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   loginForm.Username,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Printf("Failed to create the token.")
		return
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}
	http.SetCookie(w, &cookie)

	log.Printf(`| {"user": %s, "time": %d}`, loginForm.Username, time.Now().Unix())
}
