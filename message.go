package main

import "github.com/google/uuid"

type Message struct {
	ClientId uuid.UUID `json:"clientId"`
	UserId   uuid.UUID `json:userId`
	Payload  string    `json:"payload"`
}
