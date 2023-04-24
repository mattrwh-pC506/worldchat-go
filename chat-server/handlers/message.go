package handlers

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

type MessageType string

const (
	TextMessage    MessageType = "TEXT"
	SuccessMessage MessageType = "SUCCESS"
	ErrorMessage   MessageType = "ERROR"
)

type Message struct {
	Id       uuid.UUID   `json:"id"`
	ClientId uuid.UUID   `json:"clientId"`
	Type     MessageType `json:"type"`
	Payload  string      `json:"payload"`
}

func NewMessage(clientId uuid.UUID, messageType MessageType, payload string) (*Message, error) {
	if clientId == uuid.Nil {
		return nil, errors.New("invalid client ID")
	}

	if messageType != TextMessage && messageType != SuccessMessage && messageType != ErrorMessage {
		return nil, errors.New("invalid message type")
	}

	return &Message{
		Id:       uuid.New(),
		ClientId: clientId,
		Type:     messageType,
		Payload:  payload,
	}, nil
}

func NewTextMessage(clientId uuid.UUID, payload string) (*Message, error) {
	return NewMessage(clientId, TextMessage, payload)
}

func NewSuccessMessage(clientId uuid.UUID, payload string) (*Message, error) {
	return NewMessage(clientId, SuccessMessage, payload)
}

func NewErrorMessage(clientId uuid.UUID, payload string) (*Message, error) {
	return NewMessage(clientId, ErrorMessage, payload)
}

func (m *Message) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

func FromJSON(data []byte) (*Message, error) {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
