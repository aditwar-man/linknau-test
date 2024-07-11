package handlers

import (
	"encoding/json"
	"linknau-test/middleware"
	"linknau-test/models"
	"linknau-test/services"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&credentials)

	if credentials.Username == "user" && credentials.Password == "password" {
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &models.Claims{
			Username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(models.JwtKey)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		w.Write([]byte(tokenString))
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func FetchData(w http.ResponseWriter, r *http.Request) {
	url := "https://jsonplaceholder.typicode.com/todos/"
	data, err := services.FetchDataFromRemote(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

var Authenticate = middleware.Authenticate
