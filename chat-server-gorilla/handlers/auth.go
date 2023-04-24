package handlers

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateJWT(secret []byte) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ExpiresAt"] = time.Now().Add(24 * time.Hour) // expires in 24 hours
	claims["Authorized"] = true
	claims["Username"] = "n/a"

	// Sign and return the token with the secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string, secret []byte) (bool, error) {
	claims := jwt.MapClaims{}
	if _, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}); err != nil {
		return false, err
	}
	return true, nil
}
