package handlers

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Room struct {
	// public
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client

	// private
	clients     map[*Client]bool
	messages    []Message
	messageIds  map[uuid.UUID]bool
	messagesRWM sync.RWMutex
	clientsRWM  sync.RWMutex
}

func NewRoom() *Room {
	return &Room{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		messageIds: make(map[uuid.UUID]bool),
	}
}

func (room *Room) AddMessage(message Message) *Room {
	room.messagesRWM.Lock()
	defer room.messagesRWM.Unlock()

	if _, exists := room.messageIds[message.Id]; !exists {
		room.messages = append(room.messages, message)
		room.messageIds[message.Id] = true
	}

	return room
}

func (room *Room) GetMessageByIndex(index int) Message {
	room.messagesRWM.RLock()
	defer room.messagesRWM.RUnlock()

	return room.messages[index]
}

func (room *Room) GetMessageSize() int {
	room.messagesRWM.RLock()
	defer room.messagesRWM.RUnlock()

	return len(room.messages)
}

func (room *Room) Run() {
	for {
		select {
		case client := <-room.Register:
			room.clientsRWM.Lock()
			room.clients[client] = true

			for i := 0; i < room.GetMessageSize(); i++ {
				message := room.GetMessageByIndex(i)
				byteMessage, err := message.ToJSON()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Printf("Room.Register ERROR on message.ToJSON!: %v", err)
					}
					break
				}
				client.Send <- byteMessage
			}
			room.clientsRWM.Unlock()

		case client := <-room.Unregister:
			room.clientsRWM.Lock()
			if _, ok := room.clients[client]; ok {
				close(client.Send)
				delete(room.clients, client)
			}
			room.clientsRWM.Unlock()

		case message := <-room.Broadcast:
			room.clientsRWM.RLock()
			for client := range room.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(room.clients, client)
				}
			}
			room.clientsRWM.RUnlock()
		}
	}
}
