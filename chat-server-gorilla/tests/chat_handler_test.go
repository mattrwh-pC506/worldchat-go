package main_tests

import (
	"chat-server-gorilla/handlers"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestChatHandler(t *testing.T) {
	room := handlers.NewRoom()
	s := httptest.NewServer(handlers.ChatHandler(room))
	defer s.Close()
	wsURL := "ws" + strings.TrimPrefix(s.URL, "http") + "/ws"
	c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	expectedText := "hello"

	// Write a message.
	c.WriteJSON(&handlers.Message{Id: uuid.New(), Payload: expectedText})

	// Expect the server to echo the message back. Timeout if not.
	c.SetReadDeadline(time.Now().Add(time.Millisecond * 100))
	//
	//message := &handlers.Message{}
	//if err := c.ReadJSON(*message); err != nil {
	//	t.Errorf("never received a success message")
	//}
	//
	//if message.Payload != "OK" && message.Type == handlers.SuccessMessage {
	//	t.Errorf("expected '%s', got %s", expectedText, message.Payload)
	//}
}
