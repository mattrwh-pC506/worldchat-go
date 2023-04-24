package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Success bool `json:"success"`
}

type TokenCredentials struct {
	Token string `json:"token"`
}

func AuthHandler(secret []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/auth" {
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
			return
		}

		var body TokenCredentials
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if valid, err := VerifyJWT(body.Token, secret); !valid || err != nil {
			log.Println(valid, err)
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		// Create a token response object and send it as JSON
		response := SuccessResponse{Success: true}
		jsonData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
