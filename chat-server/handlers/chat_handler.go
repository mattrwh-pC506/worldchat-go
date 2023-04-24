package handlers

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Room *Room
	Conn *websocket.Conn
	Send chan []byte
	id   uuid.UUID
}

func (client *Client) readPump() {
	defer func() {
		client.Room.Unregister <- client
		client.Conn.Close()
	}()
	client.Conn.SetReadLimit(maxMessageSize)
	client.Conn.SetReadDeadline(time.Now().Add(pongWait))
	client.Conn.SetPongHandler(
		func(string) error {
			client.Conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})

	// infinite loop for reading incoming messages breaks loop on error
	for {
		_, textMessage, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Client.readPump ERROR on client.conn.ReadMessage!: %v", err)
				if errorMessage, err := NewErrorMessage(client.id, "InternalServerError", string(textMessage)); err != nil {
					client.Conn.WriteJSON(errorMessage)
				}
			}
			break
		}

		message, err := NewTextMessage(client.id, string(textMessage))
		if err != nil {
			log.Printf("Client.readPump ERROR creating NewTextMessage!: %v", err)
			if errorMessage, err := NewErrorMessage(client.id, "InternalServerError", string(textMessage)); err != nil {
				client.Conn.WriteJSON(errorMessage)
			}
			return
		}

		byteMessage, err := message.ToJSON()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Client.readPump ERROR on message.ToJSON!: %v", err)
				if errorMessage, err := NewErrorMessage(client.id, "InternalServerError", string(textMessage)); err != nil {
					client.Conn.WriteJSON(errorMessage)
				}
			}
			break
		}

		client.Room.Broadcast <- byteMessage
	}
}

func (client *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)

	// cleanup ticker and client connection on error
	defer func() {
		ticker.Stop()
		client.Conn.Close()
		errorMessage, err := NewErrorMessage(client.id, "InternalServerError", "")
		if err != nil {
			client.Conn.WriteJSON(errorMessage)
		}
	}()
	// infinite loop to handle incoming writes, returns on error
	for {
		select {
		case textMessage, ok := <-client.Send:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// send failed, send a closing message to consumer
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			}

			message, err := FromJSON(textMessage)
			if err != nil {
				return
			}
			client.Room.AddMessage(*message)

			// Send success message to client who created message, not message itself
			if message.ClientId == client.id {
				successMessage, err := NewSuccessMessage(client.id, message.Payload)
				if err != nil {
					continue
				}
				if err := client.Conn.WriteJSON(successMessage); err != nil {
					continue
				}

			} else {
				if (Message{} != *message) {
					if err := client.Conn.WriteJSON(message); err != nil {
						continue
					}
				}
			}

		case <-ticker.C:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ChatHandler(room *Room) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(responseWriter, request, nil)
		if err != nil {
			log.Printf("ChatHandler HandleFunc ERROR on upgrader.Upgrade!: %v", err)
			return
		}
		client := &Client{id: uuid.New(), Room: room, Conn: conn, Send: make(chan []byte, 256)}
		client.Room.Register <- client

		// Handle writes and reads in separate goroutines
		go client.writePump()
		go client.readPump()
	}
}
