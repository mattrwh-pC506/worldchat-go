package main_tests

import (
	"bytes"
	"chat-server/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandler(t *testing.T) {
	secret := []byte("testsecret")

	t.Run("Valid Token", func(t *testing.T) {
		token, _ := handlers.GenerateJWT(secret)

		reqBody := handlers.TokenCredentials{Token: token}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(reqBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		handler := handlers.AuthHandler(secret)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("invalid status: %d, expected: %d", rr.Code, http.StatusOK)
		}

		var response handlers.SuccessResponse
		err := json.NewDecoder(rr.Body).Decode(&response)

		if err != nil {
			t.Errorf("unknown error: %s", err)
		}

		if response.Success != true {
			t.Errorf("Auth verification unsuccessful")
		}
	})

	t.Run("Invalid Token", func(t *testing.T) {
		reqBody := handlers.TokenCredentials{Token: "invalidtoken"}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(reqBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		handler := handlers.AuthHandler(secret)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("invalid status: %d, expected: %d", rr.Code, http.StatusUnauthorized)
		}
	})
}
