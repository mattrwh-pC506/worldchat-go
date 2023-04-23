package main

import (
	"encoding/json"
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
	room *Room
	conn *websocket.Conn
	send chan []byte
	id   uuid.UUID
}

func (client *Client) readHandler() {
	defer func() {
		client.room.unregister <- client
		client.conn.Close()
	}()
	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(
		func(string) error {
			client.conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})

	// infinite loop for reading incoming messages
	// breaks loop on error
	for {
		_, textMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message := &Message{Id: uuid.New(), ClientId: client.id, Payload: string(textMessage)}
		byteMessage, err := json.Marshal(message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		client.room.broadcast <- byteMessage
	}
}

func (client *Client) writeHandler() {
	ticker := time.NewTicker(pingPeriod)

	// cleanup ticker and client connection on error
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()
	// infinite loop to handle incoming writes, returns on error
	for {
		select {
		case textMessage, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// send failed, send a closing message to consumer
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
			}

			message := Message{}
			json.Unmarshal(textMessage, &message)
			getStore().addMessage(message)

			if message.ClientId == client.id {
				continue
			}
			if (Message{} != message) {
				if err := client.conn.WriteJSON(message); err != nil {
					return
				}
			}

		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func chatHandler(room *Room, responseWriter http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{id: uuid.New(), room: room, conn: conn, send: make(chan []byte, 256)}
	client.room.register <- client

	store := getStore()
	for i := 0; i < store.getSize(); i++ {
		message := store.getMessageByIndex(i)
		byteMessage, err := json.Marshal(message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		client.room.broadcast <- byteMessage
	}

	// Handle writes and reads in separate goroutines
	go client.writeHandler()
	go client.readHandler()
}
