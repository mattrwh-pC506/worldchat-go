package main

import "github.com/google/uuid"

type Message struct {
	Id       uuid.UUID `json:"id"`
	ClientId uuid.UUID `json:"clientId"`
	UserId   uuid.UUID `json:userId`
	Payload  string    `json:"payload"`
}
