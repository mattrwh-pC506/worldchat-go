package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AuthorizeBody struct {
	Password string `json:"password"`
}
type SuccessResponse struct {
	Success bool `json:"success"`
}

func authHandler(password *string, w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/authorize" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("%v method not allowed", r.Method), http.StatusMethodNotAllowed)
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var body AuthorizeBody
	err := decoder.Decode(&body)

	if err != nil || body.Password != *password {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := SuccessResponse{Success: true}
	json.NewEncoder(w).Encode(response)
}
