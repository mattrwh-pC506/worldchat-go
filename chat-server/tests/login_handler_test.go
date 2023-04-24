package main_tests

import (
	"bytes"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"net/http/httptest"
	"testing"

	"chat-server/handlers"
)

func TestLoginHandler(t *testing.T) {
	// Set up test data and server
	password := "password123"
	secret := "secretkey123"
	handler := handlers.LoginHandler(password, []byte(secret))
	server := httptest.NewServer(handler)
	defer server.Close()

	t.Run("Test invalid credentials", func(t *testing.T) {
		reqBody := handlers.Credentials{Password: "wrongpassword"}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, server.URL+"/login", bytes.NewBuffer(reqBodyBytes))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected status %d, got %d", http.StatusUnauthorized, res.StatusCode)
		}
	})

	t.Run("Test valid credentials", func(t *testing.T) {
		// Test valid credentials
		reqBody := handlers.Credentials{Password: password}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, server.URL+"/login", bytes.NewBuffer(reqBodyBytes))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, res.StatusCode)
		}

		// Test response body contains a valid JWT token
		var tokenRes handlers.TokenResponse
		err = json.NewDecoder(res.Body).Decode(&tokenRes)
		if err != nil {
			t.Fatal(err)
		}
		_, err = jwt.Parse(tokenRes.Token, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			t.Errorf("invalid JWT token: %s", err.Error())
		}
	})
}
