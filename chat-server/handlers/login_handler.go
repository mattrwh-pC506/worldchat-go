package handlers

import (
	"encoding/json"
	"net/http"
)

type TokenResponse struct {
	Token string `json:"token"`
}

type Credentials struct {
	Password string `json:"password"`
}

func LoginHandler(password string, secret []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request contains valid credentials
		if !isValidCredential(r, password) {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		// Create a signed JWT
		tokenString, err := GenerateJWT(secret)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		// Create a token response object and send it as JSON
		response := TokenResponse{Token: tokenString}
		jsonData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func isValidCredential(r *http.Request, password string) bool {
	if r.Method != http.MethodPost || r.URL.Path != "/login" {
		return false
	}

	var body Credentials
	err := json.NewDecoder(r.Body).Decode(&body)
	return err == nil && body.Password == password
}
